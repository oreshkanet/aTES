package producer

import (
	"context"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/tasktracker"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/schema-registry"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport/mq"
)

type Producer struct {
	mq.MessageBroker
	schemaRegistry *schemaregistry.EventSchemaRegistry
	producers      map[string]mq.Producer
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

	p.producers[tasktracker.TaskStreamEvent], err = p.Produce(ctx,
		mq.NewTopic("tasktracker", tasktracker.TaskStreamEvent, "2", p.schemaRegistry))
	if err != nil {
		return err
	}

	p.producers[tasktracker.TaskAddedEvent], err = p.Produce(ctx,
		mq.NewTopic("tasktracker", tasktracker.TaskAddedEvent, "1", p.schemaRegistry))
	if err != nil {
		return err
	}

	p.producers[tasktracker.TaskAssignedEvent], err = p.Produce(ctx,
		mq.NewTopic("tasktracker", tasktracker.TaskAssignedEvent, "1", p.schemaRegistry))
	if err != nil {
		return err
	}

	p.producers[tasktracker.TaskDoneEvent], err = p.Produce(ctx,
		mq.NewTopic("tasktracker", tasktracker.TaskDoneEvent, "1", p.schemaRegistry))
	if err != nil {
		return err
	}
}
