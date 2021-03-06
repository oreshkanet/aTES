// Package service service - Имплементация бизнес-логики приложения
package service

import (
	"context"
	"github.com/oreshkanet/aTES/accounting/internal/domain"
)

type UsersService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
}

type TasksService interface {
	CreateTask(ctx context.Context, task *domain.Task) error
	UpdateTask(ctx context.Context, task *domain.Task) error
	AddTask(ctx context.Context, publicId string) error
	AssignTask(ctx context.Context, publicId string, userPublicId string) error
	DoneTask(ctx context.Context, publicId string) error
}

type AccountService interface {
	OpenTransaction(ctx context.Context, taskPublicId string, userPublicId string) error
	Payment(ctx context.Context, taskPublicId string, userPublicId string) error
	GetBalance(ctx context.Context, userPublicId string) (float32, error)
}

type Services struct {
	Users   UsersService
	Tasks   TasksService
	Account AccountService
}

type ConfigService struct {
	TasksProducer TasksEventsProducer
	AccProducer   AccountEventsProducer
	ReposUsers    UsersRepository
	ReposTasks    TasksRepository
}

func NewServices(config *ConfigService) *Services {
	return &Services{
		Users:   NewUsers(config.ReposUsers),
		Tasks:   NewTasks(config.TasksProducer, config.ReposUsers, config.ReposTasks),
		Account: NewAccount(config.AccProducer, config.ReposUsers, config.ReposTasks),
	}
}
