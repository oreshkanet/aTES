package repository

import (
	"context"
	"github.com/oreshkanet/aTES/analytics/internal/domain"
	"github.com/oreshkanet/aTES/packages/pkg/database"
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
		([public_id], [title], [assign_cost], [done_cost])
	VALUES
		(@PublicId, @Title, @AssignCost, @DoneCost)
	`

	if err := r.db.Insert(ctx, query, task); err != nil {
		return err
	}

	return nil
}
