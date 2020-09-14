package services

import (
	"github.com/edwinvautier/go-cqrs_event-sourcing/domain"
	"github.com/edwinvautier/go-cqrs_event-sourcing/forms"
	"github.com/edwinvautier/go-cqrs_event-sourcing/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) (*models.User, error) {
	var userForm forms.UserForm
	var user *models.User

	// Bind the request data to userForm type
	if err := c.ShouldBindJSON(&userForm); err != nil {
		return nil, err
	}

	// Validate user
	err := forms.ValidateUser(userForm)
	if err != nil {
		return nil, err
	}

	// Create a command and send it to the commandbus
	var command domain.CreateUserCommand
	command.Name = userForm.Name
	command.Email = userForm.Email

	u, err := domain.CommandBus.Dispatch(command)
	if err != nil {
		return nil, err
	}
	user = u.(*models.User)

	return user, nil
}
