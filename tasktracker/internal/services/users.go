package services

import "github.com/oreshkanet/aTES/tasktracker/internal/transport"

type UsersService struct {
}

func (s *UsersService) HandleUserMessage(message *transport.UserMessage) error {
	// TODO: Обработка сообщений из кафки
	return nil
}
