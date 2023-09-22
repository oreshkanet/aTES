package app

import (
	"context"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/schema-registry"
	"github.com/oreshkanet/aTES/packages/pkg/authorizer"
	"github.com/oreshkanet/aTES/packages/pkg/database"
	"github.com/oreshkanet/aTES/packages/pkg/transport/mq"
	"github.com/oreshkanet/aTES/tasktracker/internal/delivery/api"
	"github.com/oreshkanet/aTES/tasktracker/internal/delivery/events/consumer"
	"github.com/oreshkanet/aTES/tasktracker/internal/delivery/events/producer"
	"github.com/oreshkanet/aTES/tasktracker/internal/repository"
	"github.com/oreshkanet/aTES/tasktracker/internal/services"
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
}

func NewApp(
	db database.DB,
	mq mq.MessageBroker,
	http *http.Server,
	auth authorizer.AuthToken,
	schemaReg *schemaregistry.EventSchemaRegistry,
) *App {
	return &App{
		DB:        db,
		MQ:        mq,
		HTTP:      http,
		Auth:      auth,
		SchemaReg: schemaReg,
	}
}
func (a *App) Run(ctx context.Context, conf *Config) {
	// Создаём репозитории приложения
	appRepos, err := repository.NewRepository(conf.DB)
	if err != nil {
		log.Fatalf("Create repository:%s", err)
		return
	}

	// Создаём клиент для публикации нужных событий в брокера сообщений
	appEventsProducer := producer.NewProducer(conf.MQ)

	// Создаём сервисы приложения, выполняющие бизнес-логику
	appServices := services.NewServices(&services.ConfigService{
		TasksEventsProducer: appEventsProducer,
		ReposUsers:          appRepos.Users,
		ReposTasks:          appRepos.Tasks,
	})

	// Создаём консьюминг нужных событий из брокера сообщений
	appEventsConsumer := consumer.NewConsumer(appServices.Users)

	// Запускаем консьюминг и паблишинг
	err = appEventsProducer.Run(ctx)
	if err != nil {
		log.Fatalf("Create event:%s", err)
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
