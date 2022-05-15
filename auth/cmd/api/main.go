package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/oreshkanet/aTES/auth/internal/configs"
	"github.com/oreshkanet/aTES/auth/internal/server"
	"github.com/oreshkanet/aTES/auth/pkg/auth/repository"
	"github.com/oreshkanet/aTES/auth/pkg/auth/transport"
	"github.com/oreshkanet/aTES/auth/pkg/auth/usecase"
	"github.com/oreshkanet/aTES/auth/pkg/database"
)

func main() {
	config := configs.Load()

	ctx := context.Background()

	// Создаём подключение к БД
	dbURL := fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s",
		config.MsSqlUser, config.MsSqlPwd,
		config.MsSqlhost, config.MsSqlDb,
	)
	db, err := database.NewDB(ctx, dbURL)
	if err != nil {
		log.Fatalf("%s", err.Error())
		return
	}
	defer db.Close()

	// Поднимаем подключение к Кафке
	kafkaTransport := transport.CreateTransport(fmt.Sprintf("%s:%s", config.KafkaHost, config.KafkaPort))
	defer kafkaTransport.Close()

	// Создаём сервис, содержащий бизнес-логику приложения
	authUseCase := usecase.NewAuth(
		repository.CreateRepository(db),
		kafkaTransport,
		[]byte(config.SigningKey),
		config.HashSalt,
		time.Minute)

	// Запускаем http-сервер
	app := server.NewApp(authUseCase)
	if err := app.Run(config.Port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
