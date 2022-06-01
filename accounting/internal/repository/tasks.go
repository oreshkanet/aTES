package repository

import (
	"context"
	"github.com/oreshkanet/aTES/accounting/internal/domain"
	"github.com/oreshkanet/aTES/packages/pkg/database"
)

type TasksRepository struct {
	db database.DB
}

func NewTasksRepository(db database.DB) *TasksRepository {
	return &TasksRepository{
		db: db,
	}
}

func (r *TasksRepository) FindTaskByPublicId(ctx context.Context, taskID string) (*domain.Task, error) {
	task := &domain.Task{}
	query := `
	SELECT
		[public_id], [title], [description], [assign_cost], [done_cost]
	FROM
		[dbo].[tasks]
	WHERE [public_id] = @PublicId
	`

	if err := r.db.Select(ctx, task, query, database.DBParam{Name: "PublicId", Value: taskID}); err != nil {
		return task, err
	}

	return task, nil
}

func (r *TasksRepository) InsertTask(ctx context.Context, task *domain.Task) error {
	query := `
	INSERT INTO [dbo].[tasks]
		([public_id], [title], [description], [assign_cost], [done_cost])
	VALUES
		(@PublicId, @Title, @Description, 0, 0)
	`

	if err := r.db.Insert(ctx, query, task); err != nil {
		return err
	}

	return nil
}

func (r *TasksRepository) UpdateTask(ctx context.Context, task *domain.Task) error {
	query := `
	UPDATE [dbo].[tasks]
	SET
		public_id = @PublicId,
		title = @Title,
		description = @Description
	WHERE
		id = @Id
	`

	if err := r.db.Update(ctx, query, task); err != nil {
		return err
	}

	return nil
}

func (r *TasksRepository) UpdateTaskCost(ctx context.Context, task *domain.Task) error {
	query := `
	UPDATE [dbo].[tasks]
	SET
		assign_cost = @AssignCost,
		done_cost = @DoneCost
	WHERE
		id = @Id
	`

	if err := r.db.Update(ctx, query, task); err != nil {
		return err
	}

	return nil
}
