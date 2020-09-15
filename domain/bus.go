package domain

import (
	"fmt"
	"github.com/edwinvautier/go-cqrs_event-sourcing/cqrs"
)

var cmdB cqrs.CommandBus
var CommandBus = cmdB.New()

var qrB cqrs.QueryBus
var QueryBus = qrB.New()

func InitBus() {
	CommandBus.Register(CreateUserCommand{}, CreateUserCommandHandler{})
	CommandBus.Register(DeleteUserCommand{}, DeleteUserCommandHandler{})
	CommandBus.Register(EditUserCommand{}, EditUserCommandHandler{})

	QueryBus.Register(GetUsersQuery{}, GetUsersQueryHandler{})
	QueryBus.Register(GetUserByIdQuery{}, GetUserByIdQueryHandler{})
	fmt.Println("Bus ready to accept messages")
}