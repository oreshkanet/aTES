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

	// Создаём сервисы приложения, выполняющие бизнес-логику
	appServices := services.NewServices(appRepos)

	// Создаём консьюминг нужных событий из брокера сообщений
	appEvents := events.NewHandler(appServices.Users)
	err = appEvents.Init(ctx, messageBroker)
	if err != nil {
		log.Fatalf("Create events:%s", err)
		return
	}

}
