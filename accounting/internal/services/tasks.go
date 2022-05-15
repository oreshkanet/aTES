package services

import (
	"context"
	"github.com/oreshkanet/aTES/accounting/internal/domain"
)

type TasksRepository interface {
	//FindByPublicId(ctx context.Context, publicId string) (*domain.Task, error)
}

type TasksEventsProducer interface {
}

type Tasks struct {
	eventsProducer TasksEventsProducer
	reposUsers     UsersRepository
	reposTasks     TasksRepository
}

func NewTasks(events TasksEventsProducer, reposUsers UsersRepository, reposTasks TasksRepository) *Tasks {
	return &Tasks{
		eventsProducer: events,
		reposUsers:     reposUsers,
		reposTasks:     reposTasks,
	}
}

func (s *Tasks) CreateTask(ctx context.Context, task *domain.Task) error {
	// TODO: Поиск задачи в БД по publicId
	// TODO: Если задачи в БД нет, то нужно ещё её и расценить

	s.costCalculationTask(ctx, task)

	// TODO: Добавление/Обновление в БД задачи

	return nil
}

func (s *Tasks) UpdateTask(ctx context.Context, task *domain.Task) error {
	// TODO: Добавление/Обновление в БД задачи

	return nil
}

func (s *Tasks) AddTask(ctx context.Context, publicId string) error {
	// TODO: Поиск задачи в БД по publicId
	// TODO: Если в БД задачи ещё нет, то расценим её и запишем в базу

	// s.costCalculationTask(ctx, task)

	return nil
}

func (s *Tasks) costCalculationTask(ctx context.Context, task *domain.Task) error {
	// TODO: Расценить задачу
	// task.AssignCost
	// task.DoneCost
	return nil
}
