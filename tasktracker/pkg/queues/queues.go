package queues

import "context"

type Handler interface {
	HandleMessage([]byte) error
}

type Broker interface {
	Produce(ctx context.Context, topic string) error
	Consume(ctx context.Context, topic string, handler Handler) error
}

type Producer interface {
	Publish(message []byte) error
}

type Consumer interface {
	Consume() error
}
