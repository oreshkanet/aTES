package repository

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
	"github.com/oreshkanet/aTES/tasktracker/pkg/database"
	migrate "github.com/rubenv/sql-migrate"
)

type TasksRepository struct {
	db database.DB
}

func NewTasksRepository(db database.DB) (*TasksRepository, error) {
	repos := &TasksRepository{
		db: db,
	}

	if err := db.MigrateUp(repos.getMigrations()); err != nil {
		return nil, err
	}
	return repos, nil
}

func (r *TasksRepository) getMigrations() *migrate.MemoryMigrationSource {
	// FIXME: пофиксить создание автоинкрементного primary key, а ещё и уникальный индекс
	return &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			&migrate.Migration{
				Id:   "1",
				Up:   []string{"CREATE TABLE [tasks] ([id], [public_id] varchar(40), [title] varchar(250), [description] varchar(1024), [assigned_user] int)"},
				Down: []string{"DROP TABLE [tasks]"},
			},
		},
	}
}

func (r *TasksRepository) InsertTask(ctx context.Context, task *domain.Task) error {
	query := `
	INSERT INTO [tasks]
		([public_id], [title], [description], [assigned_user])
	VALUES
		(@PublicId, @Title, @Description, @AssignedUser)
	`

	if err := r.db.Insert(ctx, query, task); err != nil {
		return err
	}

	return nil
}
