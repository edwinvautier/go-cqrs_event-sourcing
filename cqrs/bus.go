package cqrs

import (
	"reflect"
	"errors"
)
type Bus interface {
	New() *Bus
	Register()
	Dispatch()
}

// Command bus ---------------------------------------
type CommandBus struct {
	handlers CommandHandlers
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

// Query bus ---------------------------------------
type QueryBus struct {
	handlers QueryHandlers
}

func (bus QueryBus) New() *QueryBus{
	return &QueryBus{
		handlers: make(QueryHandlers),
	}
}

func (bus QueryBus) Register(qr Query, handler QueryHandler) error {
	qrName := reflect.TypeOf(qr).Name()

	if reflect.TypeOf(handler.Handle).Kind() != reflect.Func {
		return errors.New("%s is not a function" + reflect.TypeOf(handler).String())
	}
	bus.handlers[qrName] = handler
	return nil
}

func (bus QueryBus) Dispatch(qr Query) (interface{}, error){
	qrName := reflect.TypeOf(qr).Name()
	return bus.handlers[qrName].Handle(qr)
}