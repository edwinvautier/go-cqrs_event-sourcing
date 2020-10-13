package domain

import (
	"github.com/edwinvautier/go-cqrs_event-sourcing/cqrs"
	"github.com/edwinvautier/go-cqrs_event-sourcing/models"
	"fmt"
	"errors"
)

type GetUsersQuery struct {}
type GetUserByIdQuery struct {
	Id uint64
}

func (q GetUsersQuery) Payload() interface{} {
	return &q
}

func (q GetUserByIdQuery) Payload() interface{} {
	return &q
}

type GetUsersQueryHandler struct {}
type GetUserByIdQueryHandler struct {}

func (ch GetUsersQueryHandler) Handle(query cqrs.Query) (interface{}, error){
	switch query.Payload().(type) {
	case *GetUsersQuery:
		// get create events
		createEvents, err := models.GetAllUserCreateEvents()
		if err != nil {
			return nil, err
		}
		// loop
		var users []models.User

		for _, event := range createEvents {
			uEvents, err := models.GetUserEventsById(event.ID)
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			u, err := models.GetUserFromEvents(uEvents)
			if err != nil {
				continue
			}
			users = append(users, u)
		}

		// return users
		return &users, nil

	default:
		return nil, errors.New("bad command type")
	}
}

func (ch GetUserByIdQueryHandler) Handle(query cqrs.Query) (interface{}, error){
	switch qr := query.Payload().(type) {
	case *GetUserByIdQuery:
		var u models.User
		userEvents, err := models.GetUserEventsById(qr.Id)
		u, err = models.GetUserFromEvents(userEvents)

		if err != nil {
			fmt.Println("Error : ", err.Error())
			return nil, err
		}
		return &u, nil

	default:
		return nil, errors.New("bad command type")
	}
}