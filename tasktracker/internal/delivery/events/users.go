package events

import (
	"context"
	"encoding/json"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
)

func (h *Handler) HandleUserStream(rawMessage []byte) error {
	var err error

	message := &domain.UserStreamMessage{}
	if err := json.Unmarshal(rawMessage, message); err != nil {
		return err
	}

	user := &domain.User{
		PublicId: message.PublicId,
		Name:     message.Name,
		Role:     message.Role,
	}

	ctx := context.Background()
	switch message.Operation {
	case "C":
		// Операция создания (обновление) пользователя в системе
		err = h.usersService.CreateUser(ctx, user)
		break
	case "U":
		// Обновление пользователя в системе
		err = h.usersService.UpdateUser(ctx, user)
		break
	case "D":
		// TODO: Удаление пользователя из системы
		break
	default:
		// TODO: Неизвестная операция
	}

	return err
}
