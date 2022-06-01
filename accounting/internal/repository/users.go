package repository

import (
	"context"
	"github.com/oreshkanet/aTES/accounting/internal/domain"
	"github.com/oreshkanet/aTES/packages/pkg/database"
)

type UsersRepository struct {
	db database.DB
}

func NewUsersRepository(db database.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (r *UsersRepository) FindUserByPublicId(ctx context.Context, userID string) (*domain.User, error) {
	user := &domain.User{}
	query := `
	SELECT
		[id], [public_id], [name], [role], [balance]
	FROM 
		[dbo].[users]
	WHERE [public_id] = @PublicId
	`

	if err := r.db.Select(ctx, user, query, database.DBParam{Name: "PublicId", Value: userID}); err != nil {
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
				([public_id], [name], [role], [balance])
			VALUES
				(@PublicId, @Name, @Role, 0)
		END
	`

	if err := r.db.Update(ctx, query, user); err != nil {
		return err
	}

	return nil
}
