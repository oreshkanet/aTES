package main

import (
	"context"
	"fmt"
	"github.com/oreshkanet/aTES/tasktracker/internal/app"
	"github.com/oreshkanet/aTES/tasktracker/internal/configs"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport/kafka"
	"github.com/oreshkanet/aTES/tasktracker/pkg/database"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	config := configs.Load()

	// Создаём подключение к БД
	dbURL := fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s",
		config.MsSqlUser, config.MsSqlPwd,
		config.MsSqlHost, config.MsSqlDb,
	)
	db, err := database.NewDBMsSQL(ctx, dbURL)
	if err != nil {
		log.Fatalf("%s", err.Error())
		return
	}
	defer db.Close()

	// Создаём подключение к брокеру сообщений Kafka
	mb := kafka.NewBrokerKafka(
		fmt.Sprintf("%s:%s", config.KafkaHost, config.KafkaPort),
		10*time.Second,
		10*time.Second,
	)
	defer mb.Close()

	application := app.NewApp()
	application.Run(ctx, db, mb)
}
