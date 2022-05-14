package events

import (
	"context"
	"encoding/json"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
)

func (p *Producer) TaskStream(ctx context.Context, message *domain.TaskStreamMessage) error {
	// TODO: Поработать с контекстом и сделать тайм-аут или вообще сделать асинхронную отправку через горутину
	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}
	p.taskStreamCh <- msg

	return err
}

func (p *Producer) TaskAdded(ctx context.Context, message *domain.TaskAddedMessage) error {
	// TODO: Поработать с контекстом и сделать тайм-аут или вообще сделать асинхронную отправку через горутину
	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}
	p.taskAddedCh <- msg

	return err
}
