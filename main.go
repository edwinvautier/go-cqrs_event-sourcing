package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/edwinvautier/go-cqrs_event-sourcing/routes"
)

func main() {
	router := gin.Default()
	routes.SetupRouter(router)

	log.Fatal(router.Run(":8080"))
}