package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	*kafka.Writer
}

func (p *Producer) PubMessage(ctx context.Context, key string, message string) error {
	return p.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key),
		Value: []byte(message),
	})
}

func NewProducer(brokerAddress string, topic string) *Producer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})

	return &Producer{writer}
}
