package models

type User struct {
	ID    uint64 `gorm:"primary_key"`
	Name  string `gorm:"size:255"`
	Email string
}

func CreateUser(user *User) (*User, error) {
	var err error

	err = db.Debug().Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func GetUsers() (*[]User, error) {
	var err error
	var users []User

	err = db.Debug().Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, nil
}

func GetUserById(id uint64) (*User, error) {
	var err error
	var user User

	err = db.Debug().First(&user, id).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

func DeleteUserById(id uint64) error {
	var err error

	err = db.Debug().Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func EditUserById(id uint64, user *User) (*User, error) {
	var err error
	user.ID = id

	err = db.Debug().Save(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}