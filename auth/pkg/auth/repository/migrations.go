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

func getMigrations() *migrate.MemoryMigrationSource {
	return &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			&migrate.Migration{
				Id:   "1",
				Up:   []string{"CREATE TABLE users (name varchar(250), password varchar(50), role varchar(50))"},
				Down: []string{"DROP TABLE users"},
			},
		},
	}
}
