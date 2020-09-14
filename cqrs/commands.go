package cqrs

type Command interface {
	Payload() interface{}
}