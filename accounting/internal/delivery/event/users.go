package event

import (
	"context"
	"encoding/json"
	"github.com/oreshkanet/aTES/accounting/internal/domain"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/auth"
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
		err = c.uSvc.CreateUser(ctx, user)
		break
	case "U":
		// Обновление пользователя в системе
		err = c.uSvc.UpdateUser(ctx, user)
		break
	case "D":
		// TODO: Удаление пользователя из системы
		break
	default:
		// TODO: Неизвестная операция
	}

	return err
}
