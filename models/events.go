package models

type UserEvent struct {
	ID 		uint64 	`gorm:"primary_key"`
	AggregateID	uint64 `gorm:"not null"`
	Name 	string	`gorm:"size:255"`
}