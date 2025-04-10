package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/muadzbayu/test-golang/config"
	"github.com/sirupsen/logrus"
)

func main() {
	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)
	db := config.NewDatabase(viperConfig, logrus.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{"Origin, Content-Type, Accept"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))

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
