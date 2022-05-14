package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
)

type TasksRepository interface {
	//FindByPublicId(ctx context.Context, publicId string) (*domain.Task, error)
}

type Tasks struct {
	reposUsers UsersRepository
	reposTasks TasksRepository
}

func NewTasks(reposUsers UsersRepository, reposTasks TasksRepository) *Tasks {
	return &Tasks{
		reposUsers: reposUsers,
		reposTasks: reposTasks,
	}
}

func (s *Tasks) AddTask(ctx context.Context, title string, description string) (*domain.Task, error) {
	task := &domain.Task{
		PublicId:    uuid.New().String(),
		Title:       title,
		Description: description,
	}

	// TODO: Добавление в базу новой задачи и получения её внутреннего ID
	// TODO: Ассайн задачи на пользователя
	// TODO: Публикация события добавления задачи в систему "tasks.added"

	return task, nil
}

func (s *Tasks) DoneTask(ctx context.Context, taskPublicId string, userPublicId string) error {
	// TODO: Поиск в базе пользователя по паблик ID
	// TODO: Поиск задачи по паблик ID
	// TODO: Проверка ассайна задачи - завершить можно только свои задачи
	// TODO: Запись в базу информации о закрытии задачи
	// TODO: Публикация события завершения задачи "tasks.done"

	return nil
}

func (s *Tasks) ReAssignTasks(ctx context.Context) error {
	// TODO: Выбираем из БД все незавершённые задачи
	// TODO: Для каждой задачи запускаем ассайн

	return nil
}
func (s *Tasks) assignTask(ctx context.Context, task *domain.Task) error {
	// TODO: Запускаем рандомайзер, определения случайного исполнителя задачи
	// TODO: Записываем в БД ассайн задачи на нового пользователя
	// TODO: Публикуем события ассайна задачи "tasks.assigned"

	return nil
}
