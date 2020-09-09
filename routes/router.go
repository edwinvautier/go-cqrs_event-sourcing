package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/edwinvautier/go-cqrs_event-sourcing/controllers"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)
}