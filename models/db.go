package models

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

// InitializeDb initialises a connection with the database
func InitializeDb(user string, password string, host string, name string, port int) {
	dbUrl := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, name)	
	
	tmpDb, err := gorm.Open("postgres", dbUrl)
	
	if err != nil {
		log.Fatal("Couldn't connect to db", err)
	}

	fmt.Println("Connected to database!")
	db = tmpDb
}

// MakeMigrations executes migrations once the db is connected
func MakeMigrations() {
	db.AutoMigrate(&User{})
}