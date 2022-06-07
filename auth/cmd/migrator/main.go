package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlserver"
	"github.com/oreshkanet/aTES/auth/internal/config"
	"github.com/oreshkanet/aTES/packages/pkg/database/mssql"
	"os"
	"os/signal"
	"syscall"
)

func run(ctx context.Context) error {
	conf := config.Load()

	// Создаём подключение к БД
	dbURL := fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s",
		conf.MsSqlUser, conf.MsSqlPwd,
		conf.MsSqlHost, "master",
	)

	db, err := mssql.NewDBMsSQL(ctx, dbURL)
	if err != nil {
		return err
	}
	defer db.Close()

	driver, err := sqlserver.WithInstance(db.DB.DB, &sqlserver.Config{})
	if err != nil {
		return fmt.Errorf("sqlserver.WithInstance: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations/mssql",
		"mssql", driver)

	ver, isDirty, err := m.Version()
	if err != nil {
		if !errors.Is(err, migrate.ErrNilVersion) {
			return fmt.Errorf("migrate.Version: %w", err)
		}
	}

	if isDirty {
		if err := m.Force(int(ver)); err != nil {
			return fmt.Errorf("migrate.Force (version %v): %w", ver, err)
		}

		if err := m.Down(); err != nil {
			return fmt.Errorf("migrate.Down (version %v): %w", ver, err)
		}
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("migrate.Up (version %v): no change", ver)
		} else {
			return fmt.Errorf("migrate.Up (version %v): %w", ver, err)
		}
	}

	return nil
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	if err := run(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
