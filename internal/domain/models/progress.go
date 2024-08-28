package models

import "time"

type Progress struct {
	ID           uint   `gorm:"primary_key"`
	UserID       uint   `gorm:"not null"`
	Language     string `gorm:"not null"`
	WordsLearned uint   `gorm:"default:0"`
	LastActive   time.Time
}
