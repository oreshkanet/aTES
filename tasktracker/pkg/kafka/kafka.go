package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type MessageHandler interface {
	HandleMessage(string) error
}

type Producer struct {
	*kafka.Conn
}

func (p *Producer) PubMessage(message []byte) error {
	_, err := p.WriteMessages(
		kafka.Message{Value: message},
	)
	return err
}

type Consumer struct {
	*kafka.Conn
}

func (c *Consumer) ConsumeMessages(handler MessageHandler) error {
	batch := c.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		err = handler.HandleMessage(string(b[:n]))
		if err != nil {
			break
		}
	}

	return batch.Close()
}

func NewProducer(ctx context.Context, brokerAddress string, topic string, partition int) (*Producer, error) {
	conn, err := kafka.DialLeader(ctx, "tcp", brokerAddress, topic, partition)
	if err != nil {
		return nil, err
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	return &Producer{conn}, nil
}

func NewConsumer(ctx context.Context, brokerAddress string, topic string, partition int) (*Consumer, error) {
	conn, err := kafka.DialLeader(ctx, "tcp", brokerAddress, topic, partition)
	if err != nil {
		return nil, err
	}
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	return &Consumer{conn}, nil
}
