package models

import (
	"fmt"
	"encoding/json"
	"errors"
)
type Event struct {
	ID 			uint64 			`gorm:"primary_key"`
	AggregateID	uint64 			`gorm:"not null"`
	Name 		string			`gorm:"size:255"`
	Payload		json.RawMessage	
}

func CreateUserEvent(user *User) error {
	var err error

	payload, err := json.Marshal(user)
	if err != nil {
		return err
	}

	createUserEvent := Event{
		Name: 			"createUser",
		AggregateID: 	0,
		Payload:		payload,
	}

	err = db.Debug().Create(&createUserEvent).Error
	if err != nil {
		return err
	}

	user.ID = createUserEvent.ID

	return nil
}

func EditUserEvent(user *User, id uint64) error {
	if UserNotExists(id) {
		return errors.New("User doesn't exist")
	}
	var err error

	payload, err := json.Marshal(user)
	if err != nil {
		return err
	}

	editUserEvent := Event{
		Name: 			"editUser",
		AggregateID: 	id,
		Payload:		payload,
	}

	err = db.Debug().Create(&editUserEvent).Error
	if err != nil {
		return err
	}
	user.ID = id

	return nil
}

func DeleteUserEvent(id uint64) error {
	var err error
	payload, err := json.Marshal(User{})

	deleteUserEvent := Event{
		Name: 			"deleteUser",
		AggregateID: 	id,
		Payload: payload,
	}

	err = db.Debug().Create(&deleteUserEvent).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUserEventsById(id uint64) ([]Event, error) {
	var err error
	var createUserEvent Event
	events := make([]Event, 0)

	// Find create event
	err = db.Debug().Where("id = ?", id).First(&createUserEvent).Error
	if err != nil {
		return events, err
	}
	events = append(events, createUserEvent)

	// Find other events ordered by id
	modifEvents := make([]Event, 0)
	err = db.Debug().Where("aggregate_id = ?", id).Find(&modifEvents).Order("id asc").Error
	if err != nil {
		return events, err
	}
	events = append(events, modifEvents...)

	return events, nil
}

type SmallEvent struct {
	ID uint64
}

func GetAllUserCreateEvents() ([]SmallEvent, error) {
	var err error
	var createEvents []SmallEvent

	err = db.Debug().Raw("SELECT id FROM events WHERE name = 'createUser'").Scan(&createEvents).Error
	if err != nil {
		return []SmallEvent{}, err
	}

	return createEvents, nil
}

func UserNotExists(id uint64) bool {
	var err error
	var event Event
	
	// check if id exists
	err = db.Debug().Where(id).First(&event).Error
	if err != nil {
		fmt.Println(err)
		return true
	}
	
	// check if deleted
	err = db.Debug().Raw("SELECT * FROM events WHERE name = 'deleteUser' AND aggregate_id = ?", id).Scan(&event).Error
	if err != nil {
		fmt.Println("no delete found : ",err)
		return false
	}

	return true
}