package transport

import (
	"context"
	"encoding/json"
	"github.com/oreshkanet/aTES/tasktracker/pkg/queues"
)

var TopicUser = "auth.user.cud.0"

// HandlerUserMessage - описание обработчика сообщений стриминга пользователей
type HandlerUserMessage interface {
	HandleUserMessage(message *UserMessage) error
}

// UserMessage - структура сообщений стриминга пользователей
type UserMessage struct {
	Operation string `json:"operation"`
	PublicId  string `json:"public_id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
}

// UsersTransport - реализация транспорта стриминга данных пользователей
type UsersTransport struct {
	handler HandlerUserMessage
}

// HandleMessage обрабатывает входящие сообщения из брокера
func (t *UsersTransport) HandleMessage(message []byte) error {
	userMessage := &UserMessage{}
	if err := json.Unmarshal(message, userMessage); err != nil {
		return err
	}
	return t.handler.HandleUserMessage(userMessage)
}

// NewUsersTransport создаёт транспорт для стриминга пользователей
func NewUsersTransport(ctx context.Context, broker queues.Broker, handler HandlerUserMessage) (*UsersTransport, error) {
	transport := &UsersTransport{
		handler: handler,
	}

	_, err := broker.Consume(ctx, TopicUser, transport)
	if err != nil {
		return nil, err
	}

	return transport, nil
}
