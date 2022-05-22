package repository

import (
	"context"
	"github.com/oreshkanet/aTES/auth/internal/domain"
	"github.com/oreshkanet/aTES/packages/pkg/database"
)

type UserRepository interface {
	MigrateUp() error
	FindUserByPublicId(ctx context.Context, id string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
}

type Repository struct {
	User UserRepository
}

func NewRepository(db database.DB) *Repository {
	return &Repository{
		User: newUser(db),
	}
}
