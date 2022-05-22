package server

import (
	"context"
	"github.com/oreshkanet/aTES/auth/internal/delivery"
	"github.com/oreshkanet/aTES/auth/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
	httpServer *http.Server

	authUseCase service.Auth
}

func NewApp(authUseCase service.Auth) *App {
	return &App{
		authUseCase: authUseCase,
	}
}

func (a *App) Run(port string) error {
	// Создаём новый роутер
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Роутим эндпойнт авторизации
	api := router.Group("/auth")
	delivery.RegisterServerEndpoits(api, a.authUseCase)

	// Задаём параметры сервера и запускаем
	a.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to start API: %+v", err)
		}
	}()

	// Обрабатываем падение сервиса
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
