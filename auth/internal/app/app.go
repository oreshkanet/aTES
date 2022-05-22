package app

import (
	"context"
	"github.com/oreshkanet/aTES/auth/internal/delivery/api"
	"github.com/oreshkanet/aTES/auth/internal/events/producer"
	"github.com/oreshkanet/aTES/auth/internal/repository"
	"github.com/oreshkanet/aTES/auth/internal/service"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/schema-registry"
	"github.com/oreshkanet/aTES/packages/pkg/authorizer"
	"github.com/oreshkanet/aTES/packages/pkg/database"
	"github.com/oreshkanet/aTES/packages/pkg/transport/mq"
	"log"
	"net/http"
)

type App struct {
}

type Config struct {
	DB        database.DB
	MQ        mq.MessageBroker
	HTTP      *http.Server
	Auth      authorizer.AuthToken
	SchemaReg *schemaregistry.EventSchemaRegistry
	HashSalt  string
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run(ctx context.Context, conf *Config) {
	// Создаём репозитории приложения
	appRepos := repository.NewRepository(conf.DB)

	// Запускаем продюсера событий
	appEventsProducer := producer.NewProducer(conf.MQ, conf.SchemaReg)
	if err := appEventsProducer.Run(ctx); err != nil {
		log.Fatalf("run event producers: %s", err.Error())
		return
	}

	// Создаём сервисы приложения, выполняющие бизнес-логику
	appServices := service.NewService(&service.ConfigService{
		Repos:     appRepos,
		Events:    appEventsProducer,
		AuthToken: conf.Auth,
		HashSalt:  conf.HashSalt,
	})

	// Запускаем API
	appAPI := api.NewApi(
		&api.Config{
			Srv:         conf.HTTP,
			Auth:        conf.Auth,
			AuthService: appServices.Auth,
		},
	)
	go func() {
		if err := appAPI.Run(); err != nil {
			log.Fatalf("Failed to start API: %+v", err)
		}
	}()
}
