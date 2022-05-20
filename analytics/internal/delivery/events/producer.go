package events

import (
	"context"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message"
	"github.com/oreshkanet/aTES/packages/pkg/transport/mq"
)

type Producer struct {
	broker mq.MessageBroker

	transactStreamCh  chan []byte
	transactPaymentCh chan []byte
}

func NewProducer(broker mq.MessageBroker) *Producer {
	client := &Producer{
		broker:            broker,
		transactStreamCh:  make(chan []byte),
		transactPaymentCh: make(chan []byte),
	}

	return client
}

func (p *Producer) Init(ctx context.Context) {
	// TODO: запускаем косьюминг топиков
	go p.broker.Produce(ctx,
		message.GetTopicName("accounting", "transact-stream", "1"),
		p.transactStreamCh)

	go p.broker.Produce(ctx,
		message.GetTopicName("accounting", "payment", "1"),
		p.transactStreamCh)
}
