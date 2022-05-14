package app

import (
	"context"
	"github.com/oreshkanet/aTES/tasktracker/internal/delivery/api"
	"github.com/oreshkanet/aTES/tasktracker/internal/delivery/events"
	"github.com/oreshkanet/aTES/tasktracker/internal/repository"
	"github.com/oreshkanet/aTES/tasktracker/internal/services"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport/mq"
	"github.com/oreshkanet/aTES/tasktracker/pkg/authorizer"
	"github.com/oreshkanet/aTES/tasktracker/pkg/database"
	"log"
	"net/http"
)

type App struct {
}

type Config struct {
	DB   database.DB
	MQ   mq.MessageBroker
	HTTP *http.Server
	Auth authorizer.AuthToken
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run(ctx context.Context, conf *Config) {
	// Создаём репозитории приложения
	appRepos, err := repository.NewRepository(conf.DB)
	if err != nil {
		log.Fatalf("Create repository:%s", err)
		return
	}

	// Создаём клиент для публикации нужных событий в брокера сообщений
	appEventsProducer := events.NewProducer(conf.MQ)

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
	err = appEventsConsumer.Init(ctx, conf.MQ)
	if err != nil {
		log.Fatalf("Create events:%s", err)
		return
	}

	// Запускаем API
	appAPI := api.NewApi(
		&api.Config{
			Srv:         conf.HTTP,
			Auth:        conf.Auth,
			TaskService: appServices.Tasks,
		},
	)
	go func() {
		if err := appAPI.Run(); err != nil {
			log.Fatalf("Failed to start API: %+v", err)
		}
	}()
}
