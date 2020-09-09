package main

import (
	// "context"
	// "log"

	// "github.com/lana/go-commandbus"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Run(":8000")
}