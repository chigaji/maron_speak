package middleware

import (
	"github.com/chigaji/maron_speak/internal/config"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleWare(cfg *config.Config) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(cfg.JWTSecret),
		// ContextKey: "user",
	})

}

func UserContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		// claims := user.Claims.(*jwt.Stan)
		userID, _ := user.Claims.GetSubject()
		c.Set("user_id", userID)
		return next(c)
	}

}
