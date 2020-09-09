package controllers

import (
	"net/http"
	"fmt"
	
	"github.com/edwinvautier/go-cqrs_event-sourcing/services"
	"github.com/edwinvautier/go-cqrs_event-sourcing/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := services.CreateUser(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "Error: " + err.Error())
		return
	}

	c.JSON(http.StatusCreated, user)
}