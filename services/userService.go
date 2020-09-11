package services

import (
	"github.com/edwinvautier/go-cqrs_event-sourcing/bus"
	cmd "github.com/edwinvautier/go-cqrs_event-sourcing/commands"
	"github.com/edwinvautier/go-cqrs_event-sourcing/forms"
	"github.com/edwinvautier/go-cqrs_event-sourcing/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) (*models.User, error) {
	var userForm forms.UserForm
	var userCreateCommand cmd.Command
	var user *models.User

	// Bind the request data to userForm type
	if err := c.ShouldBindJSON(&userForm); err != nil {
		return nil, err
	}

	err := forms.ValidateUser(userForm)
	if err != nil {
		return nil, err
	}

	// Create a command and send it to the commandbus
	userCreateCommand.Data = userForm
	userCreateCommand.Type = "UserCreate"
	
	// Send data to the bus
	userCreateCommand = bus.Send(userCreateCommand)

	user = userCreateCommand.Data.(*models.User)
	return user, nil
}
