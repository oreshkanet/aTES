package consumer

import (
	"context"
	"encoding/json"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/auth"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
)

func (c *Consumer) HandleUserStream(rawMessage []byte) error {
	var err error

	msg := new(message.EventMessage)
	if err := json.Unmarshal(rawMessage, msg); err != nil {
		return err
	}

	msgData, ok := msg.Data.(*auth.UserStreamV1)
	if !ok {
		return err
	}
	user := &domain.User{
		PublicId: msgData.PublicId,
		Name:     msgData.Name,
		Role:     msgData.Role,
	}

	ctx := context.Background()
	switch msg.EventName {
	case "Created":
		// Операция создания (обновление) пользователя в системе
		err = c.usersService.CreateUser(ctx, user)
		break
	case "Updated":
		// Обновление пользователя в системе
		err = c.usersService.UpdateUser(ctx, user)
		break
	case "Deleted":
		// TODO: Удаление пользователя из системы
		break
	default:
		// TODO: Неизвестная операция
	}

	return err
}
