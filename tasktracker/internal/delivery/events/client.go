package events

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport"
)

type Client struct {
	broker transport.MessageBroker

	taskStreamCh chan []byte
}

func NewClient(broker transport.MessageBroker) *Client {
	client := &Client{
		broker:       broker,
		taskStreamCh: make(chan []byte),
	}

	return client
}

func (c *Client) Init(ctx context.Context) {
	// TODO: запускаем косьюминг топиков
	go c.broker.Produce(ctx, domain.TaskStreamTopic, c.taskStreamCh)
}
