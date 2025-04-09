package main

import (
	"fmt"
	"log"

	"github.com/muadzbayu/test-golang/config"
	"github.com/sirupsen/logrus"
)

func main() {
	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)
	db := config.NewDatabase(viperConfig, logrus.New())

	config.AllConfig(&config.Config{
		App:    app,
		DB:     db,
		Config: viperConfig,
	})

	port := viperConfig.GetInt("APP_PORT")
	err := app.Listen(fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
