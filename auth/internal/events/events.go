package events

import (
	"context"
	"github.com/oreshkanet/aTES/auth/internal/domain"
)

type Producer interface {
	UserCreated(ctx context.Context, user *domain.User) error
	UserUpdated(ctx context.Context, user *domain.User) error
	UserRoleChanged(ctx context.Context, user *domain.User) error
}
