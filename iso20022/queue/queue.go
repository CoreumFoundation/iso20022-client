package queue

import (
	"context"
	"time"
)

type MessageQueue struct {
	sendChannel    chan []byte
	receiveChannel chan []byte
	closed         bool
}

const DefaultQueueSize = 10

func New() *MessageQueue {
	return NewWithQueueSize(DefaultQueueSize)
}

func NewWithQueueSize(queueSize int) *MessageQueue {
	return &MessageQueue{
		sendChannel:    make(chan []byte, queueSize),
		receiveChannel: make(chan []byte, queueSize),
	}
}

func (m *MessageQueue) PushToSend(msg []byte) {
	m.sendChannel <- msg
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

func (m *MessageQueue) Close() {
	if m.closed {
		return
	}
	m.closed = true
	close(m.receiveChannel)
	close(m.sendChannel)
}
