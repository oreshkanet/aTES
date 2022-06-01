package repository

import "github.com/oreshkanet/aTES/packages/pkg/database"

type Repository struct {
	Users *UsersRepository
	Tasks *TasksRepository
}

func NewRepository(db database.DB) *Repository {
	return &Repository{
		Users: NewUsersRepository(db),
		Tasks: NewTasksRepository(db),
	}
}
