package repository

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/models"
	"github.com/oreshkanet/aTES/tasktracker/pkg/database"
	migrate "github.com/rubenv/sql-migrate"
)

type UsersRepository struct {
	db database.DB
}

func (r *UsersRepository) getMigrations() *migrate.MemoryMigrationSource {
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

func (r *UsersRepository) FindUserByPublicId(ctx context.Context, userID string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT [id], [public_id], [name], [role] FROM [users] WHERE [public_id] = @PublicId
	`

	if err := r.db.Select(ctx, user, query, database.DBParam{Name: "publicId", Value: userID}); err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func (r *UsersRepository) CreateOrUpdateUser(ctx context.Context, user *models.User) error {
	query := `
	IF EXISTS (SELECT [id] FROM [users] WHERE [public_id] = @PublicId)
		BEGIN
			UPDATE [users]
			SET
				[name] = @Name
				,[role] = @Role
			WHERE [public_id] = @PublicId
		END
	ELSE
		BEGIN
			INSERT INTO [users]
				([public_id], [name], [role])
			VALUES
				(@PublicId, @Name, @Role)
		END
	`

	if err := r.db.Update(ctx, query, user); err != nil {
		return err
	}

	return nil
}

func NewUsersRepository(db database.DB) (*UsersRepository, error) {
	repos := &UsersRepository{
		db: db,
	}

	if err := db.MigrateUp(repos.getMigrations()); err != nil {
		return nil, err
	}
	return repos, nil
}
