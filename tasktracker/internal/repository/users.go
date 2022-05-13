package repository

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/models"
	"github.com/oreshkanet/aTES/tasktracker/pkg/database"
)

type UserRepository struct {
	db database.DB
}

func (r *UserRepository) FindUserByPublicId(ctx context.Context, userID string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT [id], [public_id], [name], [role] FROM [users] WHERE [public_id] = @PublicId
	`

	if err := r.db.Select(ctx, user, query, database.DBParam{Name: "publicId", Value: userID}); err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) CreateOrUpdateUser(ctx context.Context, user *models.User) error {
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
