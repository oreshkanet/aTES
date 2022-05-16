package service

import (
	"context"
	"github.com/oreshkanet/aTES/analytics/internal/domain"
)

type TasksRepository interface {
	//FindByPublicId(ctx context.Context, publicId string) (*domain.Task, error)
}

type Tasks struct {
	reposUsers UsersRepository
	reposTasks TasksRepository
}

func NewTasks(reposUsers UsersRepository, reposTasks TasksRepository) *Tasks {
	return &Tasks{
		reposUsers: reposUsers,
		reposTasks: reposTasks,
	}
}

func (s *Tasks) CreateTask(ctx context.Context, task *domain.Task) error {
	// TODO: Добавление/Обновление в БД задачи

	return nil
}

func (s *Tasks) UpdateTask(ctx context.Context, task *domain.Task) error {
	// TODO: Добавление/Обновление в БД задачи

	return nil
}
