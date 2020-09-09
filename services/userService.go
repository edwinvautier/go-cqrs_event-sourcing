package services

import (
	"errors"

	"github.com/edwinvautier/go-cqrs_event-sourcing/models"
)

func CreateUser(user *models.User) error {
	if (user.Name == "") {
		return errors.New("No username specified")
	}

	return nil
}