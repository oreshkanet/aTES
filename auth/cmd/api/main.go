package main

import (
	"log"

	"github.com/oreshkanet/aTES/auth/internal/server"
)

func main() {
	app := server.NewApp()

	if err := app.Run("8081"); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
