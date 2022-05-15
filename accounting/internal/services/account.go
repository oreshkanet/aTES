package services

import (
	"context"
)

type AccountEventsProducer interface {
}

type Account struct {
	eventsProducer AccountEventsProducer
	reposUsers     UsersRepository
	reposTasks     TasksRepository
}

func NewAccount(events TasksEventsProducer, reposUsers UsersRepository, reposTasks AccountEventsProducer) *Tasks {
	return &Tasks{
		eventsProducer: events,
		reposUsers:     reposUsers,
		reposTasks:     reposTasks,
	}
}

func (s *Tasks) DoneTask(ctx context.Context, taskPublicId string, userPublicId string) error {
	// TODO: Поиск в базе пользователя по паблик ID
	// TODO: Поиск задачи по паблик ID
	// TODO: Создание транзакции увеличения баланса на величину стоимости выполнения задачи
	// TODO: Публикация события завершения задачи "accounts.balance.increase" для записи в аудит-лог

	return nil
}

func (s *Tasks) AssignTasks(ctx context.Context, taskPublicId string, userPublicId string) error {
	// TODO: Поиск в базе пользователя по паблик ID
	// TODO: Поиск задачи по паблик ID
	// TODO: Создание транзакции уменьшения баланса на величину стоимости ассайна задачи
	// TODO: Публикация события завершения задачи "accounts.balance.decrease" для записи в аудит-лог

	return nil
}

func (s *Tasks) GetBalance(ctx context.Context, userPublicId string) (float32, error) {
	// TODO: Поиск в базе пользователя по паблик ID и получение его баланса

	return 0, nil
}
