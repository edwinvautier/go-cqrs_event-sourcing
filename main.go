package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/caarlos0/env/v6"
	"github.com/edwinvautier/go-cqrs_event-sourcing/routes"
	"github.com/edwinvautier/go-cqrs_event-sourcing/models"
	// "github.com/edwinvautier/go-cqrs_event-sourcing/bus"
)

type config struct {
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbPort     int    `env:"DB_PORT" envDefault:"5432"`
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"DB_NAME"`
}

func main() {

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	
	// Connect to database
	models.InitializeDb(cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName, cfg.DbPort)
	models.MakeMigrations()

	router := gin.Default()
	routes.SetupRouter(router)

	// bus.Init()

	log.Fatal(router.Run(":8000"))
}