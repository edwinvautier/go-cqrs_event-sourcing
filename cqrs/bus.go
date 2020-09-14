package cqrs

import (
	"reflect"
	"errors"
)

type CommandHandler interface {
	Handle(Command) (interface{}, error)
}
type CommandHandlers map[string]CommandHandler

type CommandBus struct {
	handlers CommandHandlers
}
type Bus interface {
	New() *Bus
	Register()
	Dispatch()
}

func (bus CommandBus) New() *CommandBus{
	return &CommandBus{
		handlers: make(CommandHandlers),
	}
}

func (bus CommandBus) Register(cmd Command, handler CommandHandler) error {
	cmdName := reflect.TypeOf(cmd).Name()

	if reflect.TypeOf(handler.Handle).Kind() != reflect.Func {
		return errors.New("%s is not a function" + reflect.TypeOf(handler).String())
	}
	bus.handlers[cmdName] = handler
	return nil
}

func (bus CommandBus) Dispatch(cmd Command) (interface{}, error){
	cmdName := reflect.TypeOf(cmd).Name()
	return bus.handlers[cmdName].Handle(cmd)
}