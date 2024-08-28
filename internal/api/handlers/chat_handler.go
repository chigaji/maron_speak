package handlers

import (
	"net/http"

	"github.com/chigaji/maron_speak/internal/domain/services"
	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	ChatService *services.ChatService
}

func NewChatHandler(chatSevice *services.ChatService) *ChatHandler {
	return &ChatHandler{ChatService: chatSevice}
}

func (h *ChatHandler) SaveChat(c echo.Context) error {

	username := c.Get("user_id").(uint)
	message := c.FormValue("message")
	response := c.FormValue("response")

	if err := h.ChatService.SaveChatService(username, message, response); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}

func (h *ChatHandler) GetChatHistory(c echo.Context) error {

	userID := c.Get("user_id").(uint)

	chats, err := h.ChatService.GetChatHistory(userID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, chats)
}
