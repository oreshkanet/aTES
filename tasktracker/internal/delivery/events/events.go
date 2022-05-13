package events

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
	"github.com/oreshkanet/aTES/tasktracker/internal/services"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport"
)

type Handler struct {
	broker transport.MessageBroker

	userService *services.Users
}

func newHandler(userService *services.Users) *Handler {
	handler := &Handler{
		userService: userService,
	}

	return handler
}

func (s *Handler) Init(ctx context.Context, broker transport.MessageBroker) error {
	// TODO: запускаем косьюминг топиков
	msgCh := make(chan<- []byte)
	broker.Consume(ctx, domain.UserStreamTopic, msgCh)
	go func(msgCh chan<- []byte) {
		for msg := range msgCh {
			s.HandleUserStream(msg)
		}
	}(msgCh)

}

func (s *Handler) HandleUserStream(message []byte) error {
	/*
		userMessage := &UserMessage{}
		if err := json.Unmarshal(message, userMessage); err != nil {
			return err
		}
	*/

	ctx := context.Background()
	var err error

	user := &domain.User{
		PublicId: message.PublicId,
		Name:     message.Name,
		Role:     message.Role,
	}

	switch message.Operation {
	case "C":
		// Операция создания (обновление) пользователя в системе
		err = s.userService.CreateUser(ctx, user)
		break
	case "U":
		// Обновление пользователя в системе
		err = s.userService.UpdateUser(ctx, user)
		break
	case "D":
		// TODO: Удаление пользователя из системы
		break
	default:
		// TODO: Неизвестная операция
	}

	return err
}
