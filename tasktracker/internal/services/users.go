package services

import (
	"github.com/oreshkanet/aTES/tasktracker/internal/repository"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport"
)

type UsersService struct {
	repos *repository.UserRepository
}

func (s *UsersService) HandleUserMessage(message *transport.UserMessage) error {
	//ctx := context.Background()
	switch message.Operation {
	case "C":
		// TODO: Добавление пользователя в систему
		break
	case "U":
		// TODO: Обновление пользователя в системе
		break
	case "D":
		// TODO: Удаление пользователя из системы
		break
	default:
		// TODO: Неизвестная операция
	}
	// TODO: Обработка сообщений из кафки
	return nil
}
