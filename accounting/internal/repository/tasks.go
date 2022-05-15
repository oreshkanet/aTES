package repository

import (
	"context"
	"github.com/oreshkanet/aTES/accounting/internal/domain"
	"github.com/oreshkanet/aTES/tasktracker/pkg/database"
)

type TasksRepository struct {
	db database.DB
}

func NewTasksRepository(db database.DB) (*TasksRepository, error) {
	repos := &TasksRepository{
		db: db,
	}

	return repos, nil
}

func (r *TasksRepository) InsertTask(ctx context.Context, task *domain.Task) error {
	query := `
	INSERT INTO [tasks]
		([public_id], [title], [cost])
	VALUES
		(@PublicId, @Title, @Description, @AssignedUser)
	`

	if err := r.db.Insert(ctx, query, task); err != nil {
		return err
	}

	return nil
}
