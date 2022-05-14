package repository

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
	"github.com/oreshkanet/aTES/tasktracker/pkg/database"
	migrate "github.com/rubenv/sql-migrate"
)

type UsersRepository struct {
	db database.DB
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

func (r *UsersRepository) getMigrations() *migrate.MemoryMigrationSource {
	// FIXME: пофиксить создание автоинкрементного primary key, а ещё и уникальный индекс
	return &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			&migrate.Migration{
				Id:   "1",
				Up:   []string{"CREATE TABLE [users] ([id] int, [public_id] varchar(40), [name] varchar(250), [role] varchar(50))"},
				Down: []string{"DROP TABLE [users]"},
			},
		},
	}
}

func (r *UsersRepository) FindUserByPublicId(ctx context.Context, userID string) (*domain.User, error) {
	user := &domain.User{}
	query := `
		SELECT [id], [public_id], [name], [role] FROM [users] WHERE [public_id] = @PublicId
	`

	if err := r.db.Select(ctx, user, query, database.DBParam{Name: "publicId", Value: userID}); err != nil {
		return &domain.User{}, err
	}

	return user, nil
}

func (r *UsersRepository) CreateOrUpdateUser(ctx context.Context, user *domain.User) error {
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
