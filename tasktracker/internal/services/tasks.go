package services

import (
	"github.com/google/uuid"
	"github.com/oreshkanet/aTES/tasktracker/internal/repository"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport"
)

type TasksService struct {
	reposUsers *repository.UsersRepository
	reposTasks *repository.TasksRepository
}

func (s *UsersService) HandleTaskMessage(message *transport.UserMessage) error {
	publicId := uuid.New()
}
