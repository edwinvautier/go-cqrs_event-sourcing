package models

type User struct {
	ID 		uint64 	`gorm:"primary_key"`
	Name 	string	`gorm:"size:255"`
	Email 	string
}

func CreateUser(user *User) (*User, error) {
	var err error
	
	err = db.Debug().Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}