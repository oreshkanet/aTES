package service

import (
	"context"
	"github.com/oreshkanet/aTES/analytics/internal/domain"
)

type AnalyticRepository interface {
	GetNegativeBalance(ctx context.Context) ([]*domain.User, error)
}

type Analytic struct {
	reposAnalytic AnalyticRepository
}

func NewAnalytic(reposAnalytic AnalyticRepository) *Analytic {
	return &Analytic{
		reposAnalytic: reposAnalytic,
	}
}

func (s *Analytic) GetNegativeBalance(ctx context.Context) ([]*domain.User, error) {
	return s.GetNegativeBalance(ctx)
}
