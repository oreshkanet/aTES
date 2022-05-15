package main

import (
	"context"
	"fmt"
	"github.com/oreshkanet/aTES/tasktracker/internal/app"
	"github.com/oreshkanet/aTES/tasktracker/internal/config"
	"github.com/oreshkanet/aTES/tasktracker/pkg/authorizer"
	"github.com/oreshkanet/aTES/tasktracker/pkg/database"
	"github.com/oreshkanet/aTES/tasktracker/pkg/mq/kafka"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()

	conf := config.Load()

	// Создаём подключение к БД
	dbURL := fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s",
		conf.MsSqlUser, conf.MsSqlPwd,
		conf.MsSqlHost, conf.MsSqlDb,
	)
	db, err := database.NewDBMsSQL(ctx, dbURL)
	if err != nil {
		log.Fatalf("%s", err.Error())
		return
	}
	defer db.Close()

	// Создаём подключение к брокеру сообщений Kafka
	mqBroker := kafka.NewBrokerKafka(
		fmt.Sprintf("%s:%s", conf.KafkaHost, conf.KafkaPort),
		10*time.Second,
		10*time.Second,
	)
	defer mqBroker.Close()

	httpSrv := &http.Server{
		Addr:         ":" + conf.Port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	auth := authorizer.NewJwtToken(
		conf.SigningKey,
		10*time.Minute,
	)

	application := app.NewApp()
	application.Run(ctx,
		&app.Config{
			DB:   db,
			MQ:   mqBroker,
			HTTP: httpSrv,
			Auth: auth,
		})
}
