package handlers

import (
	"net/http"
	"strconv"

	"github.com/chigaji/maron_speak/internal/domain/services"
	"github.com/labstack/echo/v4"
)

type ProgressHandler struct {
	ProgressService *services.ProgressService
}

func NewProgressHandler(progressService *services.ProgressService) *ProgressHandler {
	return &ProgressHandler{ProgressService: progressService}
}

func (h *ProgressHandler) UpdateProgress(c echo.Context) error {

	userID := c.Get("user_id").(uint)
	language := c.FormValue("language")
	wordsLearned, err := strconv.Atoi(c.FormValue("words_learned"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid words_learned value"})
	}
	if err := h.ProgressService.UpdateProgress(userID, language, uint(wordsLearned)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}

func (h *ProgressHandler) GetProgress(c echo.Context) error {

	userID := c.Get("user_id").(uint)
	language := c.QueryParam("language")

	progress, err := h.ProgressService.GetProgress(userID, language)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, progress)
}
