// Package transport: имплементация транспортного слоя приложения:
// - реализуем отправку сообщений в сервис очередей
// - отправка запросов во внешние сервисы
package transport

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/pkg/kafka"
)

var TOPIC_USER = "auth.user.cud.0"

type UsersService interface {
	HandleMessage(message string) error
}

type Transports struct {
	UserTopic *kafka.Consumer
}

func (t *Transports) Close() {
	t.UserTopic.Close()
}

func (t *Transports) UserConsume(ctx context.Context, brokerAddress string, handler UsersService) error {
	var err error
	t.UserTopic, err = kafka.NewConsumer(ctx, brokerAddress, TOPIC_USER, 0)
	if err != nil {
		return err
	}

	t.UserTopic.ConsumeMessages(handler)

	return nil
}

func NewTransport(ctx context.Context, brokerAddress string, usersService UsersService) (*Transports, error) {
	transports := &Transports{}

	if err := transports.UserConsume(ctx, brokerAddress, usersService); err != nil {
		return nil, err
	}
	return transports, nil
}
