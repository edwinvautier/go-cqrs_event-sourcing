package cqrs

type Command interface {
	Payload() interface{}
}

type CommandHandler interface {
	Handle(Command) (interface{}, error)
}
type CommandHandlers map[string]CommandHandler