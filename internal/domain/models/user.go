package models

import "time"

type User struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserResponse struct {
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
