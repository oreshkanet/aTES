package kafka

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/pkg/queues"
	"github.com/segmentio/kafka-go"
	"time"
)

type Broker struct {
	queues.Broker
	address      string
	readTimeout  time.Duration
	writeTimeout time.Duration

	producers []*Producer
	consumers []*Consumer
}

func (b *Broker) Produce(ctx context.Context, topic string) (queues.Producer, error) {
	conn, err := kafka.DialLeader(ctx, "tcp", b.address, topic, 0)
	if err != nil {
		return nil, err
	}

	producer := &Producer{conn}
	b.producers = append(b.producers, producer)

	return producer, nil
}

func (b *Broker) Consume(ctx context.Context, topic string, handler queues.Handler) (queues.Consumer, error) {
	conn, err := kafka.DialLeader(ctx, "tcp", b.address, topic, 0)
	if err != nil {
		return nil, err
	}

	consumer := &Consumer{conn}
	b.consumers = append(b.consumers, consumer)
	if err := consumer.Consume(ctx, handler); err != nil {
		return nil, err
	}
	return consumer, nil
}

func (b *Broker) Close() {
	// TODO: обработка ошибок закрытия соединений
	for _, conn := range b.producers {
		conn.Close()
	}
	for _, conn := range b.consumers {
		conn.Close()
	}
}

type Producer struct {
	*kafka.Conn
}

func (p *Producer) Publish(message []byte) error {
	_, err := p.WriteMessages(kafka.Message{
		Value: message,
	})
	if err != nil {
		return err
	}
	return nil
}

type Consumer struct {
	*kafka.Conn
}

func (c *Consumer) Consume(ctx context.Context, handler queues.Handler) error {
	batch := c.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	go func(ctx context.Context, handler queues.Handler) {
		b := make([]byte, 10e3) // 10KB max per message
		for {
			n, err := batch.Read(b)
			if err != nil {
				break
			}
			if err := handler.HandleMessage(b[:n]); err != nil {
				// TODO: при ошибке обработки сообщений пропускаем его и пишем лог
			}
		}

		if err := batch.Close(); err != nil {
			// TODO:
		}
	}(ctx, handler)
	return nil
}

func NewBrokerKafka(address string, readTimeout time.Duration, writeTimeout time.Duration) *Broker {
	return &Broker{
		address:      address,
		readTimeout:  readTimeout,
		writeTimeout: writeTimeout,
	}
}
