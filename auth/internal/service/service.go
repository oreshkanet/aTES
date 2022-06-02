package service

import (
	"context"
	"github.com/oreshkanet/aTES/auth/internal/client/event"
	"github.com/oreshkanet/aTES/auth/internal/domain"
	"github.com/oreshkanet/aTES/auth/internal/repository"
	"github.com/oreshkanet/aTES/packages/pkg/authorizer"
)

type AuthService interface {
	SignUp(ctx context.Context, user *domain.User) error
	SignIn(ctx context.Context, publicId string, pwd string) (string, error)
	ChangeRole(ctx context.Context, publicId string, role string) error
	UpdateUserProfile(ctx context.Context, user *domain.User) error
}

type Service struct {
	Auth AuthService
}

type ConfigService struct {
	Repos     *repository.Repository
	Events    event.Producer
	AuthToken authorizer.AuthToken
	HashSalt  string
}

func NewService(conf *ConfigService) *Service {
	return &Service{
		Auth: newAuth(
			conf.Repos.User,
			conf.Events,
			conf.AuthToken,
			conf.HashSalt,
		),
	}
}
