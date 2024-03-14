package event

import (
	"context"
	"github.com/oreshkanet/aTES/auth/internal/domain"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/auth"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/schema-registry"
	"github.com/oreshkanet/aTES/packages/pkg/transport/mq"
)

type EventsProducer interface {
	UserCreated(ctx context.Context, user *domain.User) error
	UserUpdated(ctx context.Context, user *domain.User) error
	UserRoleChanged(ctx context.Context, user *domain.User) error
}
type Producer struct {
	mq.MessageBroker
	sReg  *schemaregistry.EventSchemaRegistry
	prods map[string]mq.Producer
}

func NewProducer(broker mq.MessageBroker, schemaRegistry *schemaregistry.EventSchemaRegistry) *Producer {
	return &Producer{
		MessageBroker: broker,
		sReg:          schemaRegistry,
		prods:         make(map[string]mq.Producer),
	}
}

func (p *Producer) Run(ctx context.Context) error {
	var err error

	p.prods[auth.UserStreamEvent], err = p.Produce(ctx,
		mq.NewTopic("auth", auth.UserStreamEvent, "1", p.sReg))
	if err != nil {
		return err
	}

	p.prods[auth.UserRoleChangedEvent], err = p.Produce(ctx,
		mq.NewTopic("auth", auth.UserRoleChangedEvent, "1", p.sReg))
	if err != nil {
		return err
	}

	return nil
}
