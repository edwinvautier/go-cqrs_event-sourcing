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
	command := domain.CreateUserCommand{
		Name:  userForm.Name,
		Email: userForm.Email,
	}

	u, err := domain.CommandBus.Dispatch(command)
	if err != nil {
		return nil, err
	}
	user = u.(*models.User)

	return user, nil
}

func GetUsers(c *gin.Context) (*[]models.User, error) {

	var users *[]models.User

	// Create a query and send it to the commandbus
	var qr domain.GetUsersQuery

	u, err := domain.QueryBus.Dispatch(qr)
	if err != nil {
		return nil, err
	}
	users = u.(*[]models.User)

	return users, nil
}

func GetUserById(c *gin.Context) (*models.User, error) {

	var user *models.User

	// Create a query and send it to the commandbus
	qr := domain.GetUserByIdQuery{
		Id: ConvertStringToInt(c.Param("id")),
	}

	u, err := domain.QueryBus.Dispatch(qr)
	if err != nil {
		return nil, err
	}
	user = u.(*models.User)

	return user, nil
}

func EditUserById(c *gin.Context) (*models.User, error) {
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
	command := domain.EditUserCommand{
		Name:  userForm.Name,
		Email: userForm.Email,
		Id:    ConvertStringToInt(c.Param("id")),
	}

	u, err := domain.CommandBus.Dispatch(command)
	if err != nil {
		return nil, err
	}
	user = u.(*models.User)

	return user, nil
}

func DeleteUserById(c *gin.Context) error {
	// Create a command and send it to the commandbus
	command := domain.DeleteUserCommand{
		Id: ConvertStringToInt(c.Param("id")),
	}

	_, err := domain.CommandBus.Dispatch(command)
	if err != nil {
		return err
	}

	return nil
}
