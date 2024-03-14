package event

import (
	"context"
	"encoding/json"
	"github.com/oreshkanet/aTES/accounting/internal/domain"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/accounting"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/tasktracker"
)

func (p *Producer) TransactStream(ctx context.Context, trn *domain.Transaction) error {
	traceId := ctx.Value("TraceID").(string)
	taskMsg := &accounting.TransactionStreamV1{
		PublicId:     trn.PublicId,
		UserPublicId: trn.UserPublicId,
		TaskPublicId: trn.TaskPublicId,
		Credit:       trn.Credit,
		Debit:        trn.Debit,
		Status:       trn.Status,
	}
	eventMsg := message.NewEventMessage(
		traceId,
		"Created",
		"1",
		"Accounting",
		taskMsg,
	)

	msg, err := json.Marshal(eventMsg)
	if err != nil {
		return err
	}

	return p.prods[tasktracker.TaskStreamEvent].Publish(msg)
}

func (p *Producer) PaymentFinished(ctx context.Context, pay *domain.Payment) error {
	traceId := ctx.Value("TraceID").(string)
	taskMsg := &accounting.PaymentFinishedV1{
		PublicId:     pay.PublicId,
		Dt:           pay.Dt,
		UserPublicId: pay.UserPublicId,
		Amount:       pay.Amount,
	}
	eventMsg := message.NewEventMessage(
		traceId,
		"Created",
		"1",
		"Accounting",
		taskMsg,
	)

	msg, err := json.Marshal(eventMsg)
	if err != nil {
		return err
	}

	return p.prods[tasktracker.TaskStreamEvent].Publish(msg)
}
