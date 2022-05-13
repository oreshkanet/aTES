// Имплементация бизнес-логики приложения
package services

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
	"github.com/oreshkanet/aTES/tasktracker/internal/repository"
)

type UsersService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
}

type Services struct {
	Users UsersService
}

func NewServices(repos *repository.Repository) *Services {
	return &Services{
		Users: NewUsers(repos.Users),
	}
}
