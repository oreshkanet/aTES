package events

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport"
)

type Producer struct {
	broker transport.MessageBroker

	taskStreamCh chan []byte
	taskAddedCh  chan []byte
}

func NewProducer(broker transport.MessageBroker) *Producer {
	client := &Producer{
		broker:       broker,
		taskStreamCh: make(chan []byte),
		taskAddedCh:  make(chan []byte),
	}

	return client
}

func (p *Producer) Init(ctx context.Context) {
	// TODO: запускаем косьюминг топиков
	go p.broker.Produce(ctx, domain.TaskStreamTopic, p.taskStreamCh)
	go p.broker.Produce(ctx, domain.TaskAddedTopic, p.taskAddedCh)
}
