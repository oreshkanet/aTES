package repository

import "github.com/oreshkanet/aTES/tasktracker/pkg/database"

type Repository struct {
	Users *UsersRepository
}

func NewRepository(db database.DB) (*Repository, error) {
	var err error
	reposUser, err := NewUsersRepository(db)
	if err != nil {
		return nil, err
	}

	repos := &Repository{
		Users: reposUser,
	}

	return repos, nil
}
