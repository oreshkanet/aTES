package event

import (
	"context"
	"encoding/json"
	"github.com/oreshkanet/aTES/accounting/internal/domain"
)

func (c *Consumer) HandleTaskStream(rawMessage []byte) error {
	var err error

	message := &domain.TaskStreamMessage{}
	if err := json.Unmarshal(rawMessage, message); err != nil {
		return err
	}

	user := &domain.Task{
		PublicId: message.PublicId,
		Title:    message.Title,
	}

	ctx := context.Background()
	switch message.Operation {
	case "C":
		// Операция создания (обновление) пользователя в системе
		err = c.taskService.CreateTask(ctx, user)
		break
	case "U":
		// Обновление пользователя в системе
		err = c.taskService.UpdateTask(ctx, user)
		break
	case "D":
		// TODO: Удаление пользователя из системы
		break
	default:
		// TODO: Неизвестная операция
	}

	return err
}

func (c *Consumer) HandleTaskAdded(rawMessage []byte) error {
	message := &domain.TaskAddedMessage{}
	if err := json.Unmarshal(rawMessage, message); err != nil {
		return err
	}

	ctx := context.Background()
	return c.taskService.AddTask(ctx, message.PublicId)
}

func (c *Consumer) HandleTaskAssigned(rawMessage []byte) error {
	message := &domain.TaskAssignedMessage{}
	if err := json.Unmarshal(rawMessage, message); err != nil {
		return err
	}

	ctx := context.Background()
	return c.accService.AssignTasks(ctx, message.PublicId, message.UserPublicId)
}

func (c *Consumer) HandleTaskDone(rawMessage []byte) error {
	message := &domain.TaskDoneMessage{}
	if err := json.Unmarshal(rawMessage, message); err != nil {
		return err
	}

	ctx := context.Background()
	return c.accService.DoneTask(ctx, message.PublicId, message.UserPublicId)
}
