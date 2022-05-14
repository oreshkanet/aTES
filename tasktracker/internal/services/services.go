// Package services services - Имплементация бизнес-логики приложения
package services

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
)

type UsersService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
}

type TasksService interface {
	AddTask(ctx context.Context, title string, description string) (*domain.Task, error)
	DoneTask(ctx context.Context, taskPublicId string, userPublicId string) error
	ReAssignTasks(ctx context.Context) error
}

type Services struct {
	Users UsersService
	Tasks TasksService
}

type ConfigService struct {
	TasksEvents TasksEventsClient
	ReposUsers  UsersRepository
	ReposTasks  TasksRepository
}

func NewServices(config *ConfigService) *Services {
	return &Services{
		Users: NewUsers(config.ReposUsers),
		Tasks: NewTasks(config.TasksEvents, config.ReposUsers, config.ReposTasks),
	}
}
