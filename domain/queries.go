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
	// case *GetUsersQuery:
		// var u *[]models.User

		// u, err := models.GetUsers()
		// if err != nil {
		// 	fmt.Println("Error : ", err.Error())
		// 	return nil, nil
		// }
		// return u, nil

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