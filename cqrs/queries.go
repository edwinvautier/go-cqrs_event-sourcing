package cqrs

type Query interface {
	Payload() interface{}
}

type QueryHandler interface {
	Handle(Query) (interface{}, error)
}
type QueryHandlers map[string]QueryHandler