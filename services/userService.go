package services

import (
	"github.com/edwinvautier/go-cqrs_event-sourcing/forms"
	"github.com/edwinvautier/go-cqrs_event-sourcing/models"
	"errors"
	"github.com/gin-gonic/gin"
	cmd "github.com/edwinvautier/go-cqrs_event-sourcing/commands"
	"github.com/edwinvautier/go-cqrs_event-sourcing/bus"
)

func CreateUser(c *gin.Context) (models.User, error) {
	var userForm forms.UserForm
	var userCreateCommand cmd.Command
	var user models.User

	// Bind the request data to userForm type
	if err := c.ShouldBindJSON(&userForm); err != nil {
		return user, err
	}

	// Check if fields name and email exists
	if userForm.Name == "" {
		return models.User{}, errors.New("Missing field name")
	}
	if userForm.Email == "" {
		return models.User{}, errors.New("Missing field email")
	}
	
	// Create a command and send it to the commandbus
	userCreateCommand.Data = userForm
	userCreateCommand.Type = "UserCreate"
	userCreateCommand = bus.Send(userCreateCommand)

	user = userCreateCommand.Data.(models.User)
	return user,nil
}