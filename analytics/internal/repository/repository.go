package repository

import "github.com/oreshkanet/aTES/packages/pkg/database"

type Repository struct {
	Users    *UsersRepository
	Tasks    *TasksRepository
	Analytic *Analytic
}

func NewRepository(db database.DB) (*Repository, error) {
	var err error
	reposUser, err := NewUsersRepository(db)
	if err != nil {
		return nil, err
	}

	reposTasks, err := NewTasksRepository(db)
	if err != nil {
		return nil, err
	}

	reposAnalytic, err := NewAnalytic(db)
	if err != nil {
		return nil, err
	}

	repos := &Repository{
		Users:    reposUser,
		Tasks:    reposTasks,
		Analytic: reposAnalytic,
	}

	return repos, nil
}
