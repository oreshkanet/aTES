package kafka

import (
	"context"
	"github.com/oreshkanet/aTES/packages/pkg/transport/mq"
	"github.com/segmentio/kafka-go"
	"time"
)

type Broker struct {
	address      string
	readTimeout  time.Duration
	writeTimeout time.Duration

	producers []*Producer
	consumers []*Consumer
}

func NewBroker(
	address string,
	readTimeout time.Duration,
	writeTimeout time.Duration,
) *Broker {
	return &Broker{
		address:      address,
		readTimeout:  readTimeout,
		writeTimeout: writeTimeout,
		producers:    make([]*Producer, 0),
		consumers:    make([]*Consumer, 0),
	}
}

func (b *Broker) Produce(ctx context.Context, topic *mq.Topic) (mq.Producer, error) {
	conn, err := kafka.DialLeader(ctx, "tcp", b.address, topic.GetName(), 0)
	if err != nil {
		return nil, err
	}

	producer := newProducer(conn, topic)
	b.producers = append(b.producers, producer)
	producer.Run(ctx)

	return producer, nil
}

func (b *Broker) Consume(ctx context.Context, topic *mq.Topic) (mq.Consumer, error) {
	conn, err := kafka.DialLeader(ctx, "tcp", b.address, topic.GetName(), 0)
	if err != nil {
		return nil, err
	}

	consumer := newConsumer(conn, topic)
	b.consumers = append(b.consumers, consumer)
	consumer.Run(ctx)
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
