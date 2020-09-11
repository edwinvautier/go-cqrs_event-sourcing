package handlers

import (
	"github.com/edwinvautier/go-cqrs_event-sourcing/forms"
	c "github.com/edwinvautier/go-cqrs_event-sourcing/commands"
	"github.com/edwinvautier/go-cqrs_event-sourcing/models"
)

// UserCreateCommandHandler is the command handler for a new user
func UserCreateCommandHandler(cmd c.Command) *models.User {
	userForm := cmd.Data.(forms.UserForm)
	var u models.User
	u.Name = userForm.Name
	u.Email = userForm.Email
	//fmt.Println(reflect.ValueOf(cmd.Data).Type())
	models.CreateUser(&u)

	return &u
}