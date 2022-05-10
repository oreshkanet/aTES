package repository

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/models"
	"github.com/oreshkanet/aTES/tasktracker/pkg/database"
)

type UserRepository struct {
	db database.DB
}

func (r *UserRepository) FindUserById(ctx context.Context, userID string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT [id], [name], [role] FROM [users] WHERE [id] = @id
	`

	if err := r.db.Select(ctx, user, query, database.DBParam{Name: "id", Value: userID}); err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) CreateOrUpdateUser(ctx context.Context, user models.User) error {
	query := `
	IF EXISTS (SELECT [id] FROM [users] WHERE id = @id)
		BEGIN
			UPDATE [users]
			SET
				[name] = @name
				,[role] = @role
			WHERE id = @id
		END
	ELSE
		BEGIN
			INSERT INTO [users]
				([id], [name], [role])
			VALUES
				(@id, @name, @role)
		END
	`

	if err := r.db.Update(ctx, query, user); err != nil {
		return err
	}

	return nil
}
