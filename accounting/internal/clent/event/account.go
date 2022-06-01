package event

import (
	"context"
	"encoding/json"
	"github.com/oreshkanet/aTES/accounting/internal/domain"
)

func (p *event.Producer) TransactStream(ctx context.Context, message *domain.TransactStreamMessage) error {
	// TODO: Поработать с контекстом и сделать тайм-аут или вообще сделать асинхронную отправку через горутину
	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}
	p.transactStreamCh <- msg

	return err
}

func (p *event.Producer) TransactPayment(ctx context.Context, message *domain.TransactPaymentMessage) error {
	// TODO: Поработать с контекстом и сделать тайм-аут или вообще сделать асинхронную отправку через горутину
	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}
	p.transactPaymentCh <- msg

	return err
}
