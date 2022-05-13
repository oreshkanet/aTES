package transport

import "context"

type MessageBroker interface {
	Produce(ctx context.Context, topic string) (Producer, error)
	Consume(ctx context.Context, topic string, messageChannel chan<- []byte) (Consumer, error)
	Close()
}

type Producer interface {
	Publish(message []byte) error
}

type Consumer interface {
	Consume(ctx context.Context, messageChannel chan<- []byte) error
}
