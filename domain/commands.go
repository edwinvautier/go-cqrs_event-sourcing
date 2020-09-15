package domain

import (
	"github.com/edwinvautier/go-cqrs_event-sourcing/cqrs"
	"github.com/edwinvautier/go-cqrs_event-sourcing/models"
	"fmt"
	"errors"
)

type CreateUserCommand struct {
	Name string
	Email string
}
type EditUserCommand struct {
	Name string
	Email string
	Id uint64
}
type DeleteUserCommand struct {
	Id uint64
}

func (c CreateUserCommand) Payload() interface{} {
	return &c
}
func (c DeleteUserCommand) Payload() interface{} {
	return &c
}
func (c EditUserCommand) Payload() interface{} {
	return &c
}

type CreateUserCommandHandler struct {}
type DeleteUserCommandHandler struct {}
type EditUserCommandHandler struct {}

func (ch CreateUserCommandHandler) Handle(command cqrs.Command) (interface{}, error) {
	switch cmd := command.Payload().(type) {
	case *CreateUserCommand:
		u := models.User {
			Name: cmd.Name,
			Email: cmd.Email,
		}

		_, err := models.CreateUser(&u)
		if err != nil {
			fmt.Println("Error : ", err.Error())
			return nil, nil
		}
		return &u, nil
	default:
		return nil, errors.New("bad command type")
	}
}

func (ch EditUserCommandHandler) Handle(command cqrs.Command) (interface{}, error) {
	switch cmd := command.Payload().(type) {
	case *EditUserCommand:
		u := models.User {
			Name: cmd.Name,
			Email: cmd.Email,
		}

		_, err := models.EditUserById(cmd.Id, &u)
		if err != nil {
			fmt.Println("Error : ", err.Error())
			return nil, nil
		}
		return &u, nil
	default:
		return nil, errors.New("bad command type")
	}
}

func (ch DeleteUserCommandHandler) Handle(command cqrs.Command) (interface {}, error) {
	switch cmd := command.Payload().(type) {
	case *DeleteUserCommand:
		err := models.DeleteUserById(cmd.Id)
		if err != nil {
			fmt.Println("Error : ", err.Error())
			return nil, nil
		}
		return nil, nil
	default:
		return nil, errors.New("bad command type")
	}
}