package queues

import "context"

type Handler interface {
	HandleMessage([]byte) error
}

type Broker interface {
	Produce(ctx context.Context, topic string) (Producer, error)
	Consume(ctx context.Context, topic string, handler Handler) (Consumer, error)
	Close()
}

type Producer interface {
	Publish(message []byte) error
	//	Close() error
}

type Consumer interface {
	Consume(ctx context.Context, handler Handler) error
	//	Close() error
}
