package main

import (
	"github.com/oreshkanet/aTES/internal/server"
	"log"
)

func main() {
	app := server.NewApp()

	if err := app.Run("8080"); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
