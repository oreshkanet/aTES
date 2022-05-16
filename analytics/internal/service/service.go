// Package service service - Имплементация бизнес-логики приложения
package service

import (
	"context"
	"github.com/oreshkanet/aTES/analytics/internal/domain"
)

type UsersService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
}

type TasksService interface {
	CreateTask(ctx context.Context, task *domain.Task) error
	UpdateTask(ctx context.Context, task *domain.Task) error
}

type AnalyticService interface {
	GetNegativeBalance(ctx context.Context) ([]*domain.User, error)
}

type Services struct {
	Users    UsersService
	Tasks    TasksService
	Analytic AnalyticService
}

type ConfigService struct {
	ReposUsers    UsersRepository
	ReposTasks    TasksRepository
	ReposAnalytic AnalyticRepository
}

func NewServices(config *ConfigService) *Services {
	return &Services{
		Users:    NewUsers(config.ReposUsers),
		Tasks:    NewTasks(config.ReposUsers, config.ReposTasks),
		Analytic: NewAnalytic(config.ReposAnalytic),
	}
}
