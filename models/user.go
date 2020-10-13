package models

import (
	"encoding/json"
	"errors"
)

type User struct {
	ID    uint64 `gorm:"primary_key"`
	Name  string `gorm:"size:255"`
	Email string
}

func GetUserFromEvents(events []Event) (User, error) {
	var u User
	var uid uint64
	for _, event := range events {
		switch event.Name {
		case "createUser":
			json.Unmarshal(event.Payload, &u)
			uid = event.ID
			u.ID = event.ID
		case "editUser":
			json.Unmarshal(event.Payload, &u)
			u.ID = uid
		case "deleteUser":
			return u, errors.New("User deleted")
		}
	}

	return u, nil
}