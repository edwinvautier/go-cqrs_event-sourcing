package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/edwinvautier/go-cqrs_event-sourcing/controllers"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUserById)
	r.PUT("/users/:id", controllers.EditUserById)
	r.DELETE("/users/:id", controllers.DeleteUserById)
}