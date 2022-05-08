package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oreshkanet/aTES/auth/pkg/auth/delivery"
	"github.com/oreshkanet/aTES/auth/pkg/auth/usecase"
)

type App struct {
	httpServer *http.Server

	authUseCase usecase.Auth
}

func NewApp() *App {
	authUseCase := usecase.NewAuth(
		[]byte("a001c3a244ac1f9d1cc9a197cc12f9fa"),
		"affd7407a2ebab039d8fef8c6c5bbde6",
		time.Minute)
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
