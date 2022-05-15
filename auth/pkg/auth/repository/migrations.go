package repository

import (
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
)

func (r *UserRepository) MigrateUp() error {
	migrations := getMigrations()

	_, err := migrate.Exec(r.db.DB.DB, "mssql", migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("migrate DB: %w", err)
	}

	return nil
}

func getMigrations() *migrate.FileMigrationSource {
	return &migrate.FileMigrationSource{
		Dir: "repository/migrations/mssql",
	}
}
