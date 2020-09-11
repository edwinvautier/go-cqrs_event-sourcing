package bus

import (
	c "github.com/edwinvautier/go-cqrs_event-sourcing/commands"
	"github.com/edwinvautier/go-cqrs_event-sourcing/handlers"
)

// var Ch = make(chan c.Command)

// func Init() {
// 	go handlers.UserCreateCommandHandler(Ch)
// }

func Send(cmd c.Command) c.Command{
	if cmd.Type == "UserCreate" {
		cmd.Data = handlers.UserCreateCommandHandler(cmd)
		cmd.Type = "NewUser"
	}

	return cmd
}