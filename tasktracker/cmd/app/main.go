package main

import (
	"context"
	"fmt"
	"github.com/oreshkanet/aTES/tasktracker/internal/configs"
	"github.com/oreshkanet/aTES/tasktracker/internal/repository"
	"github.com/oreshkanet/aTES/tasktracker/internal/services"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport"
	"github.com/oreshkanet/aTES/tasktracker/pkg/database"
	"github.com/oreshkanet/aTES/tasktracker/pkg/queues/kafka"
	"log"
	"time"
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
	db, err := database.NewDBMsSQL(ctx, dbURL)
	if err != nil {
		log.Fatalf("%s", err.Error())
		return
	}
	defer db.Close()

	appRepos, err := repository.NewRepository(db)
	if err != nil {
		log.Fatalf("Create repository:%s", err)
		return
	}

	appServices := services.NewServices(appRepos)

	broker := kafka.NewBrokerKafka(
		fmt.Sprintf("%s:%s", config.KafkaHost, config.KafkaPort),
		10*time.Second,
		10*time.Second,
	)
	defer broker.Close()

	_, err = transport.NewTransport(ctx, broker, appServices.Users)
	if err != nil {
		log.Fatalf("start transport: %v", err)
	}
}
