package services

import (
	"github.com/chigaji/maron_speak/internal/domain/models"
	"github.com/chigaji/maron_speak/internal/repository/db"
)

type ChatService struct {
	ChatRepo *db.ChatRepo
}

func NewChatService(chatRepo *db.ChatRepo) *ChatService {
	return &ChatService{ChatRepo: chatRepo}
}

func (s *ChatService) SaveChatService(userID uint, message, response string) error {
	chat := &models.Chat{
		UserID:   userID,
		Message:  message,
		Response: response,
	}
	return s.ChatRepo.SaveChat(chat)
}

func (s *ChatService) GetChatHistory(userID uint) ([]models.Chat, error) {
	return s.ChatRepo.GetChatHistory(userID)
}
