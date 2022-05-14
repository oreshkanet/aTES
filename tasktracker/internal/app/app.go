package app

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/delivery/events"
	"github.com/oreshkanet/aTES/tasktracker/internal/repository"
	"github.com/oreshkanet/aTES/tasktracker/internal/services"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport"
	"github.com/oreshkanet/aTES/tasktracker/pkg/database"
	"log"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run(ctx context.Context, db database.DB, messageBroker transport.MessageBroker) {
	// Создаём репозитории приложения
	appRepos, err := repository.NewRepository(db)
	if err != nil {
		log.Fatalf("Create repository:%s", err)
		return
	}

	// Создаём клиент для публикации нужных событий в брокера сообщений
	appEventsProducer := events.NewProducer(messageBroker)

	// Создаём сервисы приложения, выполняющие бизнес-логику
	appServices := services.NewServices(&services.ConfigService{
		TasksEventsProducer: appEventsProducer,
		ReposUsers:          appRepos.Users,
		ReposTasks:          appRepos.Tasks,
	})

	// Создаём консьюминг нужных событий из брокера сообщений
	appEventsConsumer := events.NewConsumer(appServices.Users)

	// Запускаем консьюминг и паблишинг
	appEventsProducer.Init(ctx)
	err = appEventsConsumer.Init(ctx, messageBroker)
	if err != nil {
		log.Fatalf("Create events:%s", err)
		return
	}
}
