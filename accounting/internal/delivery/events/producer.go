package events

import (
	"context"
	"github.com/oreshkanet/aTES/accounting/internal/domain"
	"github.com/oreshkanet/aTES/accounting/internal/transport/mq"
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
	go p.broker.Produce(ctx, domain.TransactStreamTopic, p.transactStreamCh)
	go p.broker.Produce(ctx, domain.TransactPaymentTopic, p.transactStreamCh)
}
