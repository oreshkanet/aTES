package main

import (
	"context"
	"fmt"
	"github.com/oreshkanet/aTES/auth/internal/app"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/schema-registry"
	"github.com/oreshkanet/aTES/packages/pkg/authorizer"
	"log"
	"net/http"
	"time"

	"github.com/oreshkanet/aTES/auth/internal/config"
	"github.com/oreshkanet/aTES/packages/pkg/database"
	"github.com/oreshkanet/aTES/packages/pkg/transport/mq/kafka"
)

func main() {
	conf := config.Load()

	ctx := context.Background()

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

	// Поднимаем подключение к Кафке
	kafkaBroker := kafka.NewBrokerKafka(
		fmt.Sprintf("%s:%s", conf.KafkaHost, conf.KafkaPort),
		3*time.Second,
		3*time.Second,
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

	app := app.NewApp()
	app.Run(ctx, &app.Config{
		DB:        db,
		MQ:        kafkaBroker,
		HTTP:      httpSrv,
		Auth:      authToken,
		SchemaReg: schemaRegistry,
		HashSalt:  conf.HashSalt,
	})
}
