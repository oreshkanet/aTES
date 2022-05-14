package events

import (
	"context"
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
