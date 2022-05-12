// Package transport: имплементация транспортного слоя приложения:
// - реализуем отправку сообщений в сервис очередей
// - отправка запросов во внешние сервисы
package transport

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/services"
	"github.com/oreshkanet/aTES/tasktracker/pkg/queues"
)

// Transports содержит все транспорты, используемые в приложении
type Transports struct {
	Users *UsersTransport
}

// NewTransport создаёт новый экземпляр транспортов, используемых в приложении
func NewTransport(ctx context.Context, broker queues.Broker, usersService *services.UsersService) (*Transports, error) {
	var err error
	transports := &Transports{}

	transports.Users, err = NewUsersTransport(ctx, broker, usersService)
	if err != nil {
		return nil, err
	}

	return transports, nil
}
