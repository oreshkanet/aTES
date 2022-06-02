package mssql

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

type mssql struct {
	*sqlx.DB
}

func NewDBMsSQL(ctx context.Context, connURL string) (*mssql, error) {
	db, err := sqlx.ConnectContext(ctx, "mssql", connURL)
	if err != nil {
		return nil, err
	}

	return &mssql{db}, nil
}

func (d *mssql) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.GetContext(ctx, dest, query, args)
}

func (d *mssql) Insert(ctx context.Context, query string, arg interface{}) error {
	if _, err := d.NamedExecContext(ctx, query, arg); err != nil {
		return err
	}
	return nil
}

func (d *mssql) Update(ctx context.Context, query string, arg interface{}) error {
	if _, err := d.NamedExecContext(ctx, query, arg); err != nil {
		return err
	}
	return nil
}

func (d *mssql) Delete(ctx context.Context, query string, arg interface{}) error {
	if _, err := d.NamedExecContext(ctx, query, arg); err != nil {
		return err
	}
	return nil
}

func (d *mssql) MigrateUp(migrations migrate.MigrationSource) error {
	_, err := migrate.Exec(d.DB.DB, "mssql", migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("migrate DB: %w", err)
	}

	return nil
}
