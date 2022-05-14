package events

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
	"github.com/oreshkanet/aTES/tasktracker/internal/services"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport"
)

type Consumer struct {
	broker transport.MessageBroker

	usersService services.UsersService
}

func NewConsumer(usersService services.UsersService) *Consumer {
	handler := &Consumer{
		usersService: usersService,
	}

	return handler
}

func (c *Consumer) Init(ctx context.Context, broker transport.MessageBroker) error {
	// TODO: запускаем косьюминг топиков
	msgCh := make(chan []byte)
	go broker.Consume(ctx, domain.UserStreamTopic, msgCh)
	go c.Handle(ctx, msgCh, c.HandleUserStream)

	return nil
}

func (c *Consumer) Handle(ctx context.Context, msgCh <-chan []byte, handleMessages func(rawMessage []byte) error) {
	for {
		select {
		case <-ctx.Done():
			break
		case msg := <-msgCh:
			if err := handleMessages(msg); err != nil {
				// TODO: пропускать такое сообщение или пытатьсяобрабатывать снова?
				break
			}
		}
	}
}
