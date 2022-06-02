package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/oreshkanet/aTES/auth/internal/app"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/schema-registry"
	"github.com/oreshkanet/aTES/packages/pkg/authorizer"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/oreshkanet/aTES/auth/internal/config"
	"github.com/oreshkanet/aTES/packages/pkg/database"
	"github.com/oreshkanet/aTES/packages/pkg/transport/mq/kafka"
)

func run(ctx context.Context) error {
	conf := config.Load()

	// Создаём подключение к БД
	dbURL := fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s",
		conf.MsSqlUser, conf.MsSqlPwd,
		conf.MsSqlHost, conf.MsSqlDb,
	)
	db, err := database.NewDBMsSQL(ctx, dbURL)
	if err != nil {
		return err
	}

	// Поднимаем подключение к Кафке
	kafkaBroker := kafka.NewBrokerKafka(
		fmt.Sprintf("%s:%s", conf.KafkaHost, conf.KafkaPort),
		3*time.Second,
		3*time.Second,
	)

	// Подключаем регистра для схем валидации
	schemaRegistry := schemaregistry.NewRegistry(conf.SchemaRegistryPath)

	// Авторизатор сервисов SSO
	authToken := authorizer.NewJwtToken(conf.SigningKey, 10*time.Minute)

	// HTTP-сервер
	httpSrv := &http.Server{
		Addr:         ":" + conf.Port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Инстанс приложения
	appInstance := app.NewApp(db, kafkaBroker, httpSrv, authToken, schemaRegistry, conf.HashSalt)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		appInstance.Run(ctx)
		return nil
	})
	g.Go(func() error {
		<-ctx.Done()

		const timeout = 5 * time.Second
		shutdownCtx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := appInstance.Stop(shutdownCtx); err != nil {
			fmt.Printf("Failed to stop app: %s", err.Error())
		}

		return nil
	})

	if err := g.Wait(); !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	if err := run(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
