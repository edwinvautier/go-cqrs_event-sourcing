package controllers

import (
	"net/http"

	"github.com/edwinvautier/go-cqrs_event-sourcing/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	if c.Request.Header.Get("Content-Type") != "application/json" {
		c.JSON(http.StatusBadRequest, "Wrong content type, expecting application/json.")
		return
	}

	user, err := services.CreateUser(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusCreated, user)
}