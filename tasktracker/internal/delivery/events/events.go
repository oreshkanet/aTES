package events

import (
	"context"
	"encoding/json"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
	"github.com/oreshkanet/aTES/tasktracker/internal/services"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport"
)

type Handler struct {
	broker transport.MessageBroker

	usersService services.UsersService
}

func NewHandler(usersService services.UsersService) *Handler {
	handler := &Handler{
		usersService: usersService,
	}

	return handler
}

func (h *Handler) Init(ctx context.Context, broker transport.MessageBroker) error {
	// TODO: запускаем косьюминг топиков
	msgCh := make(chan []byte)
	go broker.Consume(ctx, domain.UserStreamTopic, msgCh)
	go h.Handle(ctx, msgCh, h.HandleUserStream)

	return nil
}

func (h *Handler) Handle(ctx context.Context, msgCh <-chan []byte, handleMessages func(rawMessage []byte) error) {
	for {
		select {
		case <-ctx.Done():
			break
		case msg := <-msgCh:
			if err := handleMessages(msg); err != nil {
				// TODO: пропускать такое сообщение или пытатьсяобрабатывать снова?
				break
			}
		}
	}
}

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
