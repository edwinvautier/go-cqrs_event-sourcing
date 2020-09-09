package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/edwinvautier/go-cqrs_event-sourcing/commandbus"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/users", commandbus.Send)
}