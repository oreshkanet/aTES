package kafka

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport/mq"
	"github.com/segmentio/kafka-go"
	"time"
)

type Broker struct {
	mq.MessageBroker
	address      string
	readTimeout  time.Duration
	writeTimeout time.Duration

	producers []*Producer
	consumers []*Consumer
}

func (b *Broker) Produce(ctx context.Context, topic string, messageCh <-chan []byte) error {
	conn, err := kafka.DialLeader(ctx, "tcp", b.address, topic, 0)
	if err != nil {
		return err
	}

	producer := &Producer{conn}
	b.producers = append(b.producers, producer)
	go producer.Publish(ctx, messageCh)

	return nil
}

func (b *Broker) Consume(ctx context.Context, topic string, messageCh chan<- []byte) error {
	conn, err := kafka.DialLeader(ctx, "tcp", b.address, topic, 0)
	if err != nil {
		return err
	}

	consumer := &Consumer{conn}
	b.consumers = append(b.consumers, consumer)
	consumer.Consume(ctx, messageCh)
	return nil
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

func (p *Producer) Publish(ctx context.Context, messageCh <-chan []byte) {
	for {
		select {
		case <-ctx.Done():
			p.Close()
			break
		case message := <-messageCh:
			_, err := p.WriteMessages(kafka.Message{
				Value: message,
			})
			if err != nil {
				break
			}
		}
	}
}

type Consumer struct {
	*kafka.Conn
}

func (c *Consumer) Consume(ctx context.Context, messageCh chan<- []byte) {
	batch := c.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
	b := make([]byte, 10e3)         // 10KB max per message
	for {
		select {
		case <-ctx.Done():
			batch.Close()
			break
		default:
			n, err := batch.Read(b)
			if err != nil {
				// TODO: Под вопросом. Если не смогли прочитать сообшение, возможно, будет выгоднее его пропустить
				break
			}
			// Засылаем сырые данные в канал сообщений
			messageCh <- b[:n]
		}
	}
}

func NewBrokerKafka(address string, readTimeout time.Duration, writeTimeout time.Duration) *Broker {
	return &Broker{
		address:      address,
		readTimeout:  readTimeout,
		writeTimeout: writeTimeout,
	}
}
