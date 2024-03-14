package app

import (
	"context"
	eventProducer "github.com/oreshkanet/aTES/accounting/internal/client/event"
	"github.com/oreshkanet/aTES/accounting/internal/delivery/api"
	eventConsumer "github.com/oreshkanet/aTES/accounting/internal/delivery/event"
	"github.com/oreshkanet/aTES/accounting/internal/repository"
	"github.com/oreshkanet/aTES/accounting/internal/service"
	"github.com/oreshkanet/aTES/packages/pkg/transport/mq"
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
	appRepos := repository.NewRepository(conf.DB)

	// Создаём клиент для публикации нужных событий в брокера сообщений
	appEventsProducer := eventProducer.NewProducer(conf.MQ)

	// Создаём сервисы приложения, выполняющие бизнес-логику
	appServices := service.NewServices(&service.ConfigService{
		AccProducer: appEventsProducer,
		ReposUsers:  appRepos.Users,
		ReposTasks:  appRepos.Tasks,
	})

	// Создаём консьюминг нужных событий из брокера сообщений
	appEventsConsumer := eventConsumer.NewConsumer(
		conf.MQ,
		appServices.Users,
		appServices.Tasks,
		appServices.Account,
	)

	// Запускаем консьюминг и паблишинг
	appEventsProducer.Init(ctx)
	err = appEventsConsumer.Init(ctx, conf.MQ)
	if err != nil {
		log.Fatalf("Create event:%s", err)
		return
	}

	// Запускаем API
	appAPI := api.NewApi(
		&api.Config{
			Srv:        conf.HTTP,
			Auth:       conf.Auth,
			AccService: appServices.Account,
		},
	)
	go func() {
		if err := appAPI.Run(); err != nil {
			log.Fatalf("Failed to start API: %+v", err)
		}
	}()
}
