package app

import (
	"context"
	"github.com/oreshkanet/aTES/auth/internal/client/event"
	"github.com/oreshkanet/aTES/auth/internal/delivery/api"
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
	DB        database.DB
	MQ        mq.MessageBroker
	HTTP      *http.Server
	Auth      authorizer.AuthToken
	SchemaReg *schemaregistry.EventSchemaRegistry
	HashSalt  string
}

func NewApp(
	db database.DB,
	mq mq.MessageBroker,
	http *http.Server,
	auth authorizer.AuthToken,
	schemaReg *schemaregistry.EventSchemaRegistry,
	hashSalt string) *App {
	return &App{
		DB:        db,
		MQ:        mq,
		HTTP:      http,
		Auth:      auth,
		SchemaReg: schemaReg,
		HashSalt:  hashSalt,
	}
}

func (a *App) Run(ctx context.Context) {
	// Создаём репозитории приложения
	appRepos := repository.NewRepository(a.DB)

	// Запускаем продюсера событий
	appEventsProducer := event.NewProducer(a.MQ, a.SchemaReg)
	if err := appEventsProducer.Run(ctx); err != nil {
		log.Fatalf("run event producers: %s", err.Error())
		return
	}

	// Создаём сервисы приложения, выполняющие бизнес-логику
	appServices := service.NewService(&service.ConfigService{
		Repos:     appRepos,
		Events:    appEventsProducer,
		AuthToken: a.Auth,
		HashSalt:  a.HashSalt,
	})

	// Запускаем API
	appAPI := api.NewApi(
		&api.Config{
			Srv:         a.HTTP,
			Auth:        a.Auth,
			AuthService: appServices.Auth,
		},
	)
	go func() {
		if err := appAPI.Run(); err != nil {
			log.Fatalf("Failed to start API: %+v", err)
		}
	}()
}

func (a *App) Stop(ctx context.Context) error {
	if err := a.HTTP.Shutdown(ctx); err != nil {
		return err
	}

	// TODO: разделить закрытие паблишеров и консьюмеров
	//      и таймаут, чтобы дать возможность отработать запущенные задачи
	a.MQ.Close()

	return nil
}
