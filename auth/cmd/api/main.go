package main

import (
	"log"

	"github.com/oreshkanet/aTES/auth/internal/configs"
	"github.com/oreshkanet/aTES/auth/internal/server"
)

func main() {
	config := configs.Load()
	app := server.NewApp(config.SigningKey, config.HashSalt)

	if err := app.Run(config.Port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
