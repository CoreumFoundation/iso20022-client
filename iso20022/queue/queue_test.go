package queue

import (
	"context"
	"os"
	"path"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
)

func TestSendQueue(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	requireT := require.New(t)
	ctrl := gomock.NewController(t)
	logMock := logger.NewAnyLogMock(ctrl)
	cacheDir := path.Join(os.TempDir(), "iso20022-test1")

	q := New(logMock, cacheDir)
	go func() {
		err := q.Start(ctx)
		requireT.NoError(err)
	}()
	t.Cleanup(func() {
		q.Close()
		err := os.RemoveAll(cacheDir)
		require.NoError(t, err)
	})
	message := []byte("message")

	ch := make(chan struct{})

	go func() {
		<-time.After(2 * time.Second)
		for i := 0; i < 5; i++ {
			go q.PushToSend(strconv.Itoa(i), message)
			<-time.After(599 * time.Millisecond)
		}
		close(ch)
	}()
	messages := q.PopFromSend(ctx, 11, 2*time.Second)
	requireT.Len(messages, 4)
	messages = q.PopFromSend(ctx, 1, 2*time.Second)
	requireT.Len(messages, 1)
	<-ch
	cancel()
	messages = q.PopFromSend(ctx, 1, time.Millisecond)
	requireT.Nil(messages)
}

func TestReceiveQueue(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	requireT := require.New(t)
	ctrl := gomock.NewController(t)
	logMock := logger.NewAnyLogMock(ctrl)
	cacheDir := path.Join(os.TempDir(), "iso20022-test2")

	q := New(logMock, cacheDir)
	go func() {
		err := q.Start(ctx)
		requireT.NoError(err)
	}()
	t.Cleanup(func() {
		q.Close()
		err := os.RemoveAll(cacheDir)
		require.NoError(t, err)
	})
	msg := []byte("message")

	ch := make(chan struct{})

	go func() {
		<-time.After(80 * time.Millisecond)
		q.PushToReceive(msg)
		close(ch)
	}()
	message, ok := q.PopFromReceive()
	requireT.False(ok)
	requireT.Nil(message)
	<-time.After(100 * time.Millisecond)
	message, ok = q.PopFromReceive()
	requireT.True(ok)
	requireT.NotNil(message)
	<-ch
	cancel()
	message, ok = q.PopFromReceive()
	requireT.False(ok)
	requireT.Nil(message)
}

func TestClose(t *testing.T) {
	ctrl := gomock.NewController(t)
	logMock := logger.NewAnyLogMock(ctrl)
	cacheDir := path.Join(os.TempDir(), "iso20022-test3")

	q := New(logMock, cacheDir)
	go func() {
		err := q.Start(context.Background())
		require.NoError(t, err)
	}()
	t.Cleanup(func() {
		q.Close()
		err := os.RemoveAll(cacheDir)
		require.NoError(t, err)
	})
	q.Close()
}
