package service

import (
	"context"
	"github.com/oreshkanet/aTES/accounting/internal/domain"
)

type UsersRepository interface {
	FindUserByPublicId(ctx context.Context, userID string) (*domain.User, error)
	CreateOrUpdateUser(ctx context.Context, user *domain.User) error
}

type Users struct {
	repos UsersRepository
}

func NewUsers(repos UsersRepository) *Users {
	return &Users{
		repos: repos,
	}
}

func (s *Users) CreateUser(ctx context.Context, user *domain.User) error {
	return s.repos.CreateOrUpdateUser(ctx, user)
}

func (s *Users) UpdateUser(ctx context.Context, user *domain.User) error {
	return s.repos.CreateOrUpdateUser(ctx, user)
}
