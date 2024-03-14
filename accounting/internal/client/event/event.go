package event

import (
	"context"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/accounting"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/schema-registry"
	"github.com/oreshkanet/aTES/packages/pkg/transport/mq"
)

type Producer struct {
	mq.MessageBroker
	sReg  *schemaregistry.EventSchemaRegistry
	prods map[string]mq.Producer
}

func NewProducer(
	broker mq.MessageBroker,
	schemaRegistry *schemaregistry.EventSchemaRegistry,
) *Producer {
	client := &Producer{
		MessageBroker: broker,
		sReg:          schemaRegistry,
		prods:         make(map[string]mq.Producer),
	}

	return client
}

func (p *Producer) Run(ctx context.Context) error {
	var err error

	p.prods[accounting.TransactionStreamEvent], err = p.Produce(ctx,
		mq.NewTopic("accounting", accounting.TransactionStreamEvent, "1", p.sReg))
	if err != nil {
		return err
	}

	p.prods[accounting.TaskCostCalculatedEvent], err = p.Produce(ctx,
		mq.NewTopic("accounting", accounting.TaskCostCalculatedEvent, "1", p.sReg))
	if err != nil {
		return err
	}

	p.prods[accounting.PaymentFinishedEvent], err = p.Produce(ctx,
		mq.NewTopic("accounting", accounting.PaymentFinishedEvent, "1", p.sReg))
	if err != nil {
		return err
	}

	return nil
}
