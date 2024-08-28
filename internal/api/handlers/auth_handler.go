package handlers

import (
	"fmt"
	"net/http"

	"github.com/chigaji/maron_speak/internal/domain/models"
	"github.com/chigaji/maron_speak/internal/domain/services"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Register(c echo.Context) error {
	// username := c.FormValue("username")
	// email := c.FormValue("email")
	// password := c.FormValue("password")
	userData := models.User{}
	if err := c.Bind(&userData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	fmt.Printf("User: %v", userData)

	// log.Fatalf("passed values : username: %s, email: %s, password %s", username, email, password)

	// user, err := h.AuthService.Register(userData.Username, userData.Email, userData.Password)
	user, err := h.AuthService.Register(userData.Username, userData.Email, userData.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	userResponse := models.UserResponse{Username: user.Username, Email: user.Email, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
	return c.JSON(http.StatusOK, userResponse)
}

func (h *AuthHandler) Login(c echo.Context) error {
	// username := c.FormValue("username")
	// password := c.FormValue("password")

	userData := models.User{}
	if err := c.Bind(&userData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	user, err := h.AuthService.Login(userData.Username, userData.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credential"})
	}

	userResponse := models.UserResponse{Username: user.Username, Email: user.Email, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
	return c.JSON(http.StatusOK, userResponse)
}
