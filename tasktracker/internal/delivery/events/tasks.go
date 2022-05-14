package events

import (
	"context"
	"encoding/json"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
)

func (c *Client) TaskStream(ctx context.Context, message *domain.TaskStreamMessage) error {
	// TODO: Поработать с контекстом и сделать тайм-аут или вообще сделать асинхронную отправку через горутину
	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}
	c.taskStreamCh <- msg

	return err
}

func (c *Client) TaskAdded(ctx context.Context, message *domain.TaskAddMessage) error {
	// TODO: Поработать с контекстом и сделать тайм-аут или вообще сделать асинхронную отправку через горутину
	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}
	c.taskStreamCh <- msg

	return err
}
