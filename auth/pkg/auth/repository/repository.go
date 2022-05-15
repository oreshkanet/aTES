package repository

import (
	"context"
	"database/sql"
	migrate "github.com/rubenv/sql-migrate"

	"github.com/oreshkanet/aTES/auth/pkg/auth/models"
	"github.com/oreshkanet/aTES/auth/pkg/database"
)

type UserRepository struct {
	db *database.DB
}

func (r *UserRepository) getMigrations() *migrate.MemoryMigrationSource {
	return &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			&migrate.Migration{
				Id:   "1",
				Up:   []string{"CREATE TABLE [users] ([public_id] varchar(40), [name] varchar(250), [role] varchar(50))"},
				Down: []string{"DROP TABLE [users]"},
			},
		},
	}
}

func (r *UserRepository) SelectUserByName(ctx context.Context, name string) (models.User, error) {
	var user models.User
	query := `
	SELECT name, password, role FROM users WHERE name = @name
	`
	err := r.db.GetContext(ctx, &user, query, sql.Named("name", name))
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `
    INSERT INTO users (
      name,
      password,
      role
    ) VALUES (
      :name,
      :password,
      :role
    );
  `
	if _, err := r.db.NamedExecContext(ctx, query, user); err != nil {
		return err
	}
	return nil
}

func CreateRepository(db *database.DB) *UserRepository {
	repos := &UserRepository{
		db: db,
	}

	repos.MigrateUp()
	return repos
}
