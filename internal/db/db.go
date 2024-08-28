package db

import (
	"fmt"

	"github.com/chigaji/maron_speak/internal/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {

	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)

	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return db, nil

}
