package events

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport/mq"
)

type Producer struct {
	broker mq.MessageBroker

	taskStreamCh   chan []byte
	taskAddedCh    chan []byte
	taskAssignedCh chan []byte
	taskDoneCh     chan []byte
}

func NewProducer(broker mq.MessageBroker) *Producer {
	client := &Producer{
		broker:         broker,
		taskStreamCh:   make(chan []byte),
		taskAddedCh:    make(chan []byte),
		taskAssignedCh: make(chan []byte),
		taskDoneCh:     make(chan []byte),
	}

	return client
}

func (p *Producer) Init(ctx context.Context) {
	// TODO: запускаем косьюминг топиков
	go p.broker.Produce(ctx, domain.TaskStreamTopic, p.taskStreamCh)
	go p.broker.Produce(ctx, domain.TaskAddedTopic, p.taskAddedCh)
	go p.broker.Produce(ctx, domain.TaskAssignedTopic, p.taskAssignedCh)
	go p.broker.Produce(ctx, domain.TaskDoneTopic, p.taskDoneCh)
}
