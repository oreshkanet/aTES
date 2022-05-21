package kafka

import (
	"context"
	"github.com/oreshkanet/aTES/packages/pkg/transport/mq"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	*kafka.Conn
	topic  *mq.Topic
	valid  mq.Validator
	msgCh  chan []byte
	readCh chan interface{}
}

func newConsumer(conn *kafka.Conn, topic *mq.Topic, valid mq.Validator) *Consumer {
	return &Consumer{
		Conn:  conn,
		topic: topic,
		valid: valid,
	}
}

func (c *Consumer) Run(ctx context.Context) {
	go func() {
		// Ждём начала чтения
	f1:
		for {
			select {
			case <-c.readCh:
				break f1
			default:
				continue
			}
		}

		// Запускаем чтение из коннекта брокера
		batch := c.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
		b := make([]byte, 10e3)         // 10KB max per message
	f2:
		for {
			select {
			case <-ctx.Done():
				batch.Close()
				break f2
			default:
				n, err := batch.Read(b)
				if err != nil {
					// TODO: Ошибка чтения - подумать политику ретрая (переподъём подключения или падение сервиса)
				}
				msg := b[:n]
				// Валидируем сообщение
				err = c.valid.ValidateBytes(msg, c.topic.Domain, c.topic.Event, c.topic.Version)
				if err != nil {
					// Пропускаем невалидные сообщения
					// TODO: Логируем проблемное сообщение
					continue
				}

				// Засылаем сырые данные в канал сообщений
				c.msgCh <- msg
			}
		}
	}()
}

func (c *Consumer) Read() <-chan []byte {
	if c.msgCh == nil {
		c.msgCh = make(chan []byte)
		c.readCh <- 0 // Отправляем в канал данные, чтобы запустить консьюминг топика
	}
	return c.msgCh
}
