package domain

import (
	"github.com/edwinvautier/go-cqrs_event-sourcing/cqrs"
	"github.com/edwinvautier/go-cqrs_event-sourcing/models"
	"fmt"
)

type CreateUserCommand struct {
	Name string
	Email string
}

func (c CreateUserCommand) Payload() interface{} {
	return &c
}

// User create command handler
type CreateUserCommandHandler struct {}

func (ch CreateUserCommandHandler) Handle(cmd cqrs.Command) (interface{}, error){
	cmdUser := cmd.Payload().(*CreateUserCommand)
	
	var u models.User
	u.Name = cmdUser.Name
	u.Email = cmdUser.Email

	_, err := models.CreateUser(&u)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return nil, nil
	}

	return &u, nil
}