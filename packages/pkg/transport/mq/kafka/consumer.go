package kafka

import (
	"context"
	"github.com/oreshkanet/aTES/packages/pkg/transport/mq"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	*kafka.Conn
	topic  *mq.Topic
	msgCh  chan []byte
	readCh chan interface{}
}

func newConsumer(conn *kafka.Conn, topic *mq.Topic) *Consumer {
	return &Consumer{
		Conn:  conn,
		topic: topic,
		msgCh: nil,
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
					// TODO: Под вопросом. Если не смогли прочитать сообшение, возможно, будет выгоднее его пропустить
					break f2
				}
				// Засылаем сырые данные в канал сообщений
				c.msgCh <- b[:n]
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
