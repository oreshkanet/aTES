package kafka

import (
	"context"
	"fmt"
	"github.com/oreshkanet/aTES/packages/pkg/transport/mq"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	*kafka.Conn
	topic *mq.Topic
	valid mq.Validator
	msgCh chan []byte
	// TODO: Политика ретрая публикации
}

func newProducer(conn *kafka.Conn, topic *mq.Topic, valid mq.Validator) *Producer {
	return &Producer{
		Conn:  conn,
		topic: topic,
		valid: valid,
		msgCh: nil,
	}
}

func (p *Producer) Run(ctx context.Context) {
	go func() {
		p.msgCh = make(chan []byte)

		for {
			select {
			case <-ctx.Done():
				p.Close()
				break
			case message := <-p.msgCh:
				_, err := p.WriteMessages(kafka.Message{
					Value: message,
				})
				if err != nil {
					// TODO: Реализовать настраиваемую политику ретрая для событий
					break
				}
			}
		}
	}()
}

func (p *Producer) Publish(msg []byte) error {
	// Валидируем сообщение
	err := p.valid.ValidateBytes(msg, p.topic.Domain, p.topic.Event, p.topic.Version)
	if err != nil {
		return err
	}

	if p.msgCh == nil {
		return fmt.Errorf("producer %s not run", p.topic.GetName())
	}
	// TODO: Таймаут публикации и политика ретрая для топика
	p.msgCh <- msg
	return nil
}
