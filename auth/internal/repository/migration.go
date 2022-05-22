package repository

import (
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
)

func (r *user) MigrateUp() error {
	migrations := getMigrations()

	err := r.db.MigrateUp(migrations)
	if err != nil {
		return fmt.Errorf("migrate DB: %w", err)
	}

	return nil
}

func getMigrations() *migrate.FileMigrationSource {
	return &migrate.FileMigrationSource{
		Dir: "repository/migration/mssql",
	}
}
