package domain

import (
	"fmt"
	"github.com/edwinvautier/go-cqrs_event-sourcing/cqrs"
)

var cmdB cqrs.CommandBus
var CommandBus = cmdB.New()

func InitBus() {
	CommandBus.Register(CreateUserCommand{}, CreateUserCommandHandler{})
	fmt.Println("Bus ready to accept messages")
}
