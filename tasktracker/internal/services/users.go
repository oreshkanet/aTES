package services

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/models"
	"github.com/oreshkanet/aTES/tasktracker/internal/repository"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport"
)

type UsersService struct {
	repos *repository.UserRepository
}

func (s *UsersService) HandleUserMessage(message *transport.UserMessage) error {
	ctx := context.Background()
	var err error

	user := &models.User{
		PublicId: message.PublicId,
		Name:     message.Name,
		Role:     message.Role,
	}

	switch message.Operation {
	case "C":
		// Операция создания (обновление) пользователя в системе
		err = s.repos.CreateOrUpdateUser(ctx, user)
		break
	case "U":
		// Обновление пользователя в системе
		err = s.repos.CreateOrUpdateUser(ctx, user)
		break
	case "D":
		// TODO: Удаление пользователя из системы
		break
	default:
		// TODO: Неизвестная операция
	}

	return err
}
