package models

import "time"

type Chat struct {
	ID        uint   `gorm:"primary_key"`
	UserID    uint   `gorm:"not null"`
	Message   string `gorm:"not null"`
	Response  string `gorm:"not null"`
	CreatedAt time.Time
}
