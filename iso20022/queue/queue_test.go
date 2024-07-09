package queue

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSendQueue(t *testing.T) {
	requireT := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	q := New()
	t.Cleanup(q.Close)
	message := []byte("message")

	ch := make(chan struct{})

	go func() {
		<-time.After(2 * time.Second)
		for i := 0; i < 20; i++ {
			q.PushToSend(message)
			<-time.After(99 * time.Millisecond)
		}
		close(ch)
	}()
	messages := q.PopFromSend(ctx, 11, time.Second)
	requireT.Len(messages, 10)
	messages = q.PopFromSend(ctx, 10, time.Second)
	requireT.Len(messages, 10)
	<-ch
	cancel()
	messages = q.PopFromSend(ctx, 1, time.Millisecond)
	requireT.Nil(messages)
}

func TestReceiveQueue(t *testing.T) {
	requireT := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	q := New()
	t.Cleanup(q.Close)
	msg := []byte("message")

	ch := make(chan struct{})

	go func() {
		<-time.After(80 * time.Millisecond)
		q.PushToReceive(msg)
		close(ch)
	}()
	message, ok := q.PopFromReceive(ctx, 50*time.Millisecond)
	requireT.False(ok)
	requireT.Nil(message)
	message, ok = q.PopFromReceive(ctx, 50*time.Millisecond)
	requireT.True(ok)
	requireT.NotNil(message)
	<-ch
	cancel()
	message, ok = q.PopFromReceive(ctx, 50*time.Millisecond)
	requireT.False(ok)
	requireT.Nil(message)
}

func TestClose(t *testing.T) {
	q := New()
	t.Cleanup(q.Close)
	q.Close()
}
