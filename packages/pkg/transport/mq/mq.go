package mq

import (
	"context"
)

type MessageBroker interface {
	Produce(ctx context.Context, topic *Topic) (Producer, error)
	Consume(ctx context.Context, topic *Topic) (Consumer, error)
	Close()
}

type Producer interface {
	Run(ctx context.Context)
	Publish(message []byte) error
}

type Consumer interface {
	Run(ctx context.Context)
	Read() <-chan []byte
}
