package event

import (
	"context"
	"encoding/json"
	"github.com/oreshkanet/aTES/accounting/internal/domain"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/tasktracker"
)

func (c *Consumer) HandleTaskStream(rawMessage []byte) error {
	var err error

	msg := new(message.EventMessage)
	if err = json.Unmarshal(rawMessage, msg); err != nil {
		return err
	}

	msgData, ok := msg.Data.(*tasktracker.TaskStreamMessageV2)
	if !ok {
		return err
	}
	task := &domain.Task{
		PublicId:    msgData.PublicId,
		Title:       msgData.Title,
		Description: msgData.Description,
	}

	ctx := context.Background()
	switch msg.EventName {
	case "Created":
		// Операция создания (обновление) пользователя в системе
		err = c.tSvc.CreateTask(ctx, task)
		break
	case "Updated":
		// Обновление пользователя в системе
		err = c.tSvc.UpdateTask(ctx, task)
		break
	case "Deleted":
		// TODO: Удаление пользователя из системы
		break
	default:
		// TODO: Неизвестная операция
	}

	return err
}

func (c *Consumer) HandleTaskAdded(rawMessage []byte) error {
	var err error

	msg := new(message.EventMessage)
	if err = json.Unmarshal(rawMessage, msg); err != nil {
		return err
	}

	msgData, ok := msg.Data.(*tasktracker.TaskAddedMessageV1)
	if !ok {
		return err
	}

	ctx := context.Background()
	return c.tSvc.AddTask(ctx, msgData.PublicId)
}

func (c *Consumer) HandleTaskAssigned(rawMessage []byte) error {
	var err error

	msg := new(message.EventMessage)
	if err = json.Unmarshal(rawMessage, msg); err != nil {
		return err
	}

	msgData, ok := msg.Data.(*tasktracker.TaskAssignedMessageV1)
	if !ok {
		return err
	}

	ctx := context.Background()
	return c.tSvc.AssignTask(ctx, msgData.PublicId, msgData.UserPublicId)
}

func (c *Consumer) HandleTaskDone(rawMessage []byte) error {
	var err error

	msg := new(message.EventMessage)
	if err = json.Unmarshal(rawMessage, msg); err != nil {
		return err
	}

	msgData, ok := msg.Data.(*tasktracker.TaskDoneMessageV1)
	if !ok {
		return err
	}

	ctx := context.Background()
	return c.tSvc.DoneTask(ctx, msgData.PublicId)
}
