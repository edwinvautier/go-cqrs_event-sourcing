package forms

import (
	"github.com/asaskevich/govalidator"
)

type UserForm struct {
	Name  string `valid:"stringlength(3|50)"`
	Email string `valid:"email"`
}

func ValidateUser(user UserForm) error{
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}
