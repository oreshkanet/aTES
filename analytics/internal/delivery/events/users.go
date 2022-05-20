package events

import (
	"context"
	"encoding/json"
	"github.com/oreshkanet/aTES/analytics/internal/domain"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/auth"
)

func (c *Consumer) HandleUserStream(rawMessage []byte) error {
	var err error

	msg := &message.EventMessage{}
	if err := json.Unmarshal(rawMessage, msg); err != nil {
		return err
	}
	msgData := msg.Data.(*auth.UserStreamV1)

	user := &domain.User{
		PublicId: msgData.PublicId,
		Name:     msgData.Name,
		Role:     msgData.Role,
	}

	ctx := context.Background()
	switch msgData.Operation {
	case "C":
		// Операция создания (обновление) пользователя в системе
		err = c.usersService.CreateUser(ctx, user)
		break
	case "U":
		// Обновление пользователя в системе
		err = c.usersService.UpdateUser(ctx, user)
		break
	case "D":
		// TODO: Удаление пользователя из системы
		break
	default:
		// TODO: Неизвестная операция
	}

	return err
}
