package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func NewDB(ctx context.Context, connURL string) (*DB, error) {
	db, err := sqlx.ConnectContext(ctx, "mssql", connURL)
	if err != nil {
		return nil, err
	}

	return &DB{
		DB: db,
	}, nil
}
