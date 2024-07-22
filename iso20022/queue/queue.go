package queue

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sync/atomic"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"go.uber.org/zap"

	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
)

type Status string

const (
	StatusError   Status = "error"
	StatusSending Status = "sending"
	StatusSent    Status = "sent"
)

type MessageQueue struct {
	log            logger.Logger
	cacheDir       string
	sendChannel    chan []byte
	receiveChannel chan []byte
	closed         atomic.Bool
	statuses       *ttlcache.Cache[string, Status]
}

const DefaultQueueSize = 50

func New(log logger.Logger, cacheDir string) *MessageQueue {
	return NewWithQueueSize(log, cacheDir, DefaultQueueSize)
}

func NewWithQueueSize(log logger.Logger, cacheDir string, queueSize int) *MessageQueue {
	return NewWithQueueSizeAndCacheDur(log, cacheDir, queueSize, time.Hour)
}

func NewWithQueueSizeAndCacheDur(log logger.Logger, cacheDir string, queueSize int, cacheDuration time.Duration) *MessageQueue {
	ctx := context.Background()

	if err := os.MkdirAll(cacheDir, 0777); err != nil {
		log.Error(ctx, "could not make the cache dir", zap.String("path", cacheDir), zap.Error(err))
	}

	stateFile := path.Join(cacheDir, "state.json")
	if stat, err := os.Stat(stateFile); os.IsNotExist(err) || stat.Size() == 0 {
		err = os.WriteFile(stateFile, []byte("[]"), 0777)
		if err != nil {
			log.Error(ctx, "could not write the state file", zap.String("path", stateFile), zap.Error(err))
		}
	}

	return &MessageQueue{
		log:            log,
		cacheDir:       cacheDir,
		sendChannel:    make(chan []byte, queueSize),
		receiveChannel: make(chan []byte, queueSize),
		statuses: ttlcache.New[string, Status](
			ttlcache.WithTTL[string, Status](cacheDuration),
		),
	}
}

func (m *MessageQueue) Start(ctx context.Context) error {
	m.statuses.OnEviction(m.onEviction)
	m.load(ctx)
	// TODO: We can have a goroutine here to save the state periodically
	m.statuses.Start()
	<-ctx.Done()
	m.Close()
	m.log.Info(context.Background(), "saved state before exiting the program")
	return nil
}

func (m *MessageQueue) idToFilePath(id string) string {
	h := md5.Sum([]byte(id))
	fileName := fmt.Sprintf("%x.xml", h)
	return path.Join(m.cacheDir, fileName)
}

func (m *MessageQueue) PushToSend(id string, msg []byte) Status {
	item := m.statuses.Get(id)
	if item == nil || item.Value() == StatusError {
		filePath := m.idToFilePath(id)
		if err := os.WriteFile(filePath, msg, 0777); err != nil {
			m.log.Error(
				context.Background(),
				"could not cache message",
				zap.String("messageID", id),
				zap.Error(err),
			)
		}
		m.statuses.Set(id, StatusSending, ttlcache.NoTTL)
		m.sendChannel <- msg
		return StatusSending
	}
	return item.Value()
}

func (m *MessageQueue) PushToReceive(msg []byte) {
	m.receiveChannel <- msg
}

func (m *MessageQueue) PopFromSend(ctx context.Context, n int, dur time.Duration) [][]byte {
	res := make([][]byte, 0, n)
	timer := time.NewTimer(dur)
	for {
		select {
		case msg := <-m.sendChannel:
			res = append(res, msg)
			if len(res) == n {
				return res
			}
		case <-timer.C:
			if len(res) == 0 {
				timer.Reset(dur)
				continue
			}
			return res
		case <-ctx.Done():
			return nil
		}
	}
}

func (m *MessageQueue) PopFromReceive() ([]byte, bool) {
	select {
	case msg := <-m.receiveChannel:
		return msg, true
	default:
		return nil, false
	}
}

func (m *MessageQueue) SetStatus(id string, status Status) {
	switch status {
	case StatusSending:
		m.log.Warn(context.Background(), "cannot set status to sending")
	case StatusSent:
		filePath := m.idToFilePath(id)
		err := os.Remove(filePath)
		if err != nil {
			m.log.Warn(
				context.Background(),
				"could not remove cache file",
				zap.String("path", filePath),
				zap.Error(err),
			)
		}
		m.statuses.Set(id, StatusSent, ttlcache.DefaultTTL)
	case StatusError:
		m.statuses.Set(id, StatusError, ttlcache.DefaultTTL)
	}
}

func (m *MessageQueue) Close() {
	if !m.closed.CompareAndSwap(false, true) {
		return
	}
	m.statuses.Stop()
	m.save()
	close(m.receiveChannel)
	close(m.sendChannel)
}

type cacheItem struct {
	Key       string `json:"k"`
	Status    Status `json:"s"`
	ExpiresAt int64  `json:"e"`
}

func (m *MessageQueue) onEviction(_ context.Context, reason ttlcache.EvictionReason, item *ttlcache.Item[string, Status]) {
	if reason != ttlcache.EvictionReasonExpired {
		return
	}

	filePath := m.idToFilePath(item.Key())
	err := os.Remove(filePath)
	if err != nil {
		m.log.Warn(
			context.Background(),
			"could not remove cache file",
			zap.String("path", filePath),
			zap.Error(err),
		)
	}
}

func (m *MessageQueue) load(ctx context.Context) {
	stateFilePath := path.Join(m.cacheDir, "state.json")
	data, err := os.ReadFile(stateFilePath)
	if err != nil {
		m.log.Error(context.Background(), "could not load the state", zap.String("path", stateFilePath), zap.Error(err))
		return
	}

	var items []cacheItem
	err = json.Unmarshal(data, &items)
	if err != nil {
		m.log.Warn(context.Background(), "could not unmarshal the state. starting fresh", zap.String("path", stateFilePath), zap.Error(err))
		return
	}

	sendingMessages := make([]string, 0)
	for _, item := range items {
		if item.Status == StatusSending {
			sendingMessages = append(sendingMessages, item.Key)
			continue
		}

		ttl := time.Until(time.UnixMicro(item.ExpiresAt))
		if ttl <= 0 {
			ttl = 10 * time.Millisecond
		}

		m.statuses.Set(item.Key, item.Status, ttl)
	}

	<-time.After(10 * time.Millisecond)
	m.statuses.DeleteExpired()

	for _, id := range sendingMessages {
		id := id
		filePath := m.idToFilePath(id)
		message, err := os.ReadFile(filePath)
		if err != nil {
			m.log.Error(
				ctx,
				"could not load a cached message",
				zap.String("messageID", id),
				zap.Error(err),
			)
			continue
		}
		go m.PushToSend(id, message)
	}
}

func (m *MessageQueue) save() {
	items := make([]cacheItem, 0, m.statuses.Len())
	m.statuses.Range(func(item *ttlcache.Item[string, Status]) bool {
		items = append(items, cacheItem{
			Key:       item.Key(),
			Status:    item.Value(),
			ExpiresAt: item.ExpiresAt().UnixMicro(),
		})
		return true
	})
	data, err := json.Marshal(items)
	if err != nil {
		m.log.Error(context.Background(), "could not marshal the state", zap.Error(err))
		return
	}

	err = os.WriteFile(path.Join(m.cacheDir, "state.json"), data, 0777)
	if err != nil {
		m.log.Error(context.Background(), "could not save the state", zap.Error(err))
	}
}
