package repository

import (
	"context"
	"github.com/oreshkanet/aTES/analytics/internal/domain"
	"github.com/oreshkanet/aTES/packages/pkg/database"
)

type Analytic struct {
	db database.DB
}

func NewAnalytic(db database.DB) (*Analytic, error) {
	repos := &Analytic{
		db: db,
	}

	return repos, nil
}

func (r *Analytic) GetNegativeBalance(ctx context.Context) ([]*domain.User, error) {
	query := `
	SELECT
		[public_id], [name], [balance]
	FROM [dbo].[users]
	WHERE [balance] < 0
	`
	users := make([]*domain.User, 0)
	if err := r.db.Select(ctx, users, query, nil); err != nil {
		return nil, err
	}

	return users, nil
}
