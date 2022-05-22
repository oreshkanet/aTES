package main

import (
	"context"
	"fmt"
	schemaregistry "github.com/oreshkanet/aTES/event-registry/go/pkg/schema-registry"
	"github.com/oreshkanet/aTES/tasktracker/internal/app"
	"github.com/oreshkanet/aTES/tasktracker/internal/config"
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
	kafkaBroker := kafka.NewBrokerKafka(
		fmt.Sprintf("%s:%s", conf.KafkaHost, conf.KafkaPort),
		10*time.Second,
		10*time.Second,
	)
	defer kafkaBroker.Close()

	// Подключаем регистра для схем валидации
	schemaRegistry := schemaregistry.NewRegistry(conf.SchemaRegistryPath)

	authToken := authorizer.NewJwtToken(conf.SigningKey, 10*time.Minute)

	httpSrv := &http.Server{
		Addr:         ":" + conf.Port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app := app.NewApp(db, kafkaBroker, httpSrv, authToken, schemaRegistry)
	app.Run(ctx)
}
