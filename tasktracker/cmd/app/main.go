package main

import (
	"context"
	"fmt"
	"github.com/oreshkanet/aTES/tasktracker/internal/configs"
	"github.com/oreshkanet/aTES/tasktracker/internal/services"
	"github.com/oreshkanet/aTES/tasktracker/internal/transport"
	"github.com/oreshkanet/aTES/tasktracker/pkg/database"
	"log"
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

	appServices := services.NewServices()

	appTransport, err := transport.NewTransport(ctx,
		fmt.Sprintf("%s:%s", config.KafkaHost, config.KafkaPort),
		appServices.Users)
	if err != nil {
		log.Fatalf("start transport: %v", err)
	}
	defer appTransport.Close()
}
