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
	switch cmd.Type {
	case "UserCreate":
		cmd.Data = handlers.UserCreateCommandHandler(cmd)
	}

	return cmd
}