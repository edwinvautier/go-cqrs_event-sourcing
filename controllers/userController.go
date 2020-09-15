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

func GetUsers(c *gin.Context) {
	users, err := services.GetUsers(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	user, err := services.GetUserById(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUserById(c *gin.Context) {
	err := services.DeleteUserById(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, "user successfully deleted")
}

func EditUserById(c *gin.Context) {

	if c.Request.Header.Get("Content-Type") != "application/json" {
		c.JSON(http.StatusBadRequest, "Wrong content type, expecting application/json.")
		return
	}

	user, err := services.EditUserById(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}