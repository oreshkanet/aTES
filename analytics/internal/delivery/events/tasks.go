package events

import (
	"context"
	"encoding/json"
	"github.com/oreshkanet/aTES/analytics/internal/domain"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/tasktracker"
)

func (c *Consumer) HandleTaskStream(rawMessage []byte) error {
	var err error

	msg := &message.EventMessage{}
	if err := json.Unmarshal(rawMessage, msg); err != nil {
		return err
	}
	msgData := msg.Data.(*tasktracker.TaskStreamMessageV1)

	task := &domain.Task{
		PublicId: msgData.PublicId,
		Title:    msgData.Title,
	}

	ctx := context.Background()
	switch msgData.Operation {
	case "C":
		// Операция создания (обновление) пользователя в системе
		err = c.taskService.CreateTask(ctx, task)
		break
	case "U":
		// Обновление пользователя в системе
		err = c.taskService.UpdateTask(ctx, task)
		break
	case "D":
		// TODO: Удаление пользователя из системы
		break
	default:
		// TODO: Неизвестная операция
	}

	return err
}
