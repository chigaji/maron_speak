package db

import (
	"github.com/chigaji/maron_speak/internal/domain/models"
	"github.com/jinzhu/gorm"
)

type ChatRepo struct {
	DB *gorm.DB
}

func NewChartRepo(db *gorm.DB) *ChatRepo {
	return &ChatRepo{DB: db}
}

func (r *ChatRepo) SaveChat(chart *models.Chat) error {
	return r.DB.Create(chart).Error
}

func (r *ChatRepo) GetChatHistory(userID uint) ([]models.Chat, error) {
	var chats []models.Chat
	err := r.DB.Where("user_id = ?", userID).Find(&chats).Error
	return chats, err
}
