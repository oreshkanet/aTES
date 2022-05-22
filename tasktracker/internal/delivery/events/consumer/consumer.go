package consumer

import (
	"context"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/auth"
	schemaregistry "github.com/oreshkanet/aTES/event-registry/go/pkg/schema-registry"
	"github.com/oreshkanet/aTES/tasktracker/internal/services"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport/mq"
)

type Consumer struct {
	mq.MessageBroker
	schemaRegistry *schemaregistry.EventSchemaRegistry
	consumers      map[string]mq.Consumer

	usersService services.UsersService
}

func NewConsumer(
	broker mq.MessageBroker,
	schemaRegistry *schemaregistry.EventSchemaRegistry,
	usersService services.UsersService,
) *Consumer {
	return &Consumer{
		MessageBroker:  broker,
		schemaRegistry: schemaRegistry,
		consumers:      make(map[string]mq.Consumer),
		usersService:   usersService,
	}
}

func (c *Consumer) Run(ctx context.Context) error {
	var err error
	c.consumers[auth.UserStreamEvent], err = c.Consume(ctx,
		mq.NewTopic("auth", auth.UserStreamEvent, "1", c.schemaRegistry))
	if err != nil {
		return err
	}
	go c.Handle(ctx, auth.UserStreamEvent, c.HandleUserStream)

	return nil
}

func (c *Consumer) Handle(ctx context.Context, eventName string, handler func(rawMessage []byte) error) {
	for {
		select {
		case <-ctx.Done():
			break
		case msg := <-c.consumers[eventName].Read():
			if err := handler(msg); err != nil {
				// TODO: пропускать такое сообщение или пытатьсяобрабатывать снова?
				break
			}
		}
	}
}
