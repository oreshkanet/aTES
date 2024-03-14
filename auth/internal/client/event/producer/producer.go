package producer

import (
	"context"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/auth"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/schema-registry"
	"github.com/oreshkanet/aTES/packages/pkg/transport/mq"
)

type Producer struct {
	mq.MessageBroker
	schemaRegistry *schemaregistry.EventSchemaRegistry

	producers map[string]mq.Producer
}

func NewProducer(broker mq.MessageBroker, schemaRegistry *schemaregistry.EventSchemaRegistry) *Producer {
	return &Producer{
		MessageBroker:  broker,
		schemaRegistry: schemaRegistry,
		producers:      make(map[string]mq.Producer),
	}
}

func (p *Producer) Run(ctx context.Context) error {
	var err error

	p.producers[auth.UserStreamEvent], err = p.Produce(ctx,
		mq.NewTopic("auth", auth.UserStreamEvent, "1", p.schemaRegistry))
	if err != nil {
		return err
	}

	p.producers[auth.UserRoleChangedEvent], err = p.Produce(ctx,
		mq.NewTopic("auth", auth.UserRoleChangedEvent, "1", p.schemaRegistry))
	if err != nil {
		return err
	}

	return nil
}
