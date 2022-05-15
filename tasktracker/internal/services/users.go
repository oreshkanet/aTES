package services

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
)

type UsersRepository interface {
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

// TODO: ChangeRoleUser(ctx context.Context, publicId string, role string) error
