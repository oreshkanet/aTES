package repository

import (
	"context"
	"github.com/oreshkanet/aTES/auth/internal/domain"
	"github.com/oreshkanet/aTES/packages/pkg/database"
)

type user struct {
	db database.DB
}

func newUser(db database.DB) *user {
	return &user{
		db: db,
	}
}

func (r *user) FindUserByPublicId(ctx context.Context, id string) (*domain.User, error) {
	user := &domain.User{}
	query := `
		SELECT [id], [public_id], [name], [password], [role] FROM [users] WHERE [public_id] = @PublicId
	`

	if err := r.db.Select(ctx, user, query, database.DBParam{Name: "PublicId", Value: id}); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *user) CreateUser(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO [users]
			([public_id], [name], [password], [role])
		VALUES
			(@PublicId, @Name, @Password, @Role)
	`
	if err := r.db.Insert(ctx, query, user); err != nil {
		return err
	}
	return nil
}

func (r *user) UpdateUser(ctx context.Context, user *domain.User) error {
	query := `
		UPDATE [users]
		SET
			[name] = @Name
			,[role] = @Role
			,[password] = @Password
		WHERE [id] = @Id
  `
	if err := r.db.Update(ctx, query, user); err != nil {
		return err
	}
	return nil
}
