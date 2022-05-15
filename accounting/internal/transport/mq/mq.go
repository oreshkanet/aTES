package mq

//TODO: перевезти в отдельный пакет packages

import "context"

type MessageBroker interface {
	Produce(ctx context.Context, topic string, messageChannel <-chan []byte) error
	Consume(ctx context.Context, topic string, messageChannel chan<- []byte) error
	Close()
}

type Producer interface {
	Publish(message []byte, messageChannel <-chan []byte)
}

type Consumer interface {
	Consume(ctx context.Context, messageChannel chan<- []byte)
}
