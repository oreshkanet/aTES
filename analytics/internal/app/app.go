package app

import (
	"context"
	"github.com/oreshkanet/aTES/analytics/internal/delivery/api"
	"github.com/oreshkanet/aTES/analytics/internal/delivery/events"
	"github.com/oreshkanet/aTES/analytics/internal/repository"
	"github.com/oreshkanet/aTES/analytics/internal/service"
	"github.com/oreshkanet/aTES/packages/pkg/authorizer"
	"github.com/oreshkanet/aTES/packages/pkg/database"
	"github.com/oreshkanet/aTES/packages/pkg/transport/mq"
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

	// Создаём сервисы приложения, выполняющие бизнес-логику
	appServices := service.NewServices(&service.ConfigService{
		ReposUsers:    appRepos.Users,
		ReposTasks:    appRepos.Tasks,
		ReposAnalytic: appRepos.Analytic,
	})

	// Создаём консьюминг нужных событий из брокера сообщений
	appEventsConsumer := events.NewConsumer(
		appServices.Users,
		appServices.Tasks,
		appServices.Analytic,
	)

	// Запускаем консьюминг и паблишинг
	err = appEventsConsumer.Init(ctx, conf.MQ)
	if err != nil {
		log.Fatalf("Create events:%s", err)
		return
	}

	// Запускаем API
	appAPI := api.NewApi(
		&api.Config{
			Srv:      conf.HTTP,
			Auth:     conf.Auth,
			Analytic: appServices.Analytic,
		},
	)
	go func() {
		if err := appAPI.Run(); err != nil {
			log.Fatalf("Failed to start API: %+v", err)
		}
	}()
}
