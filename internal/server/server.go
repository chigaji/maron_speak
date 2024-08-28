package server

import (
	"github.com/chigaji/maron_speak/internal/api/handlers"
	"github.com/chigaji/maron_speak/internal/config"
	"github.com/chigaji/maron_speak/internal/db"
	"github.com/chigaji/maron_speak/internal/domain/services"
	repo "github.com/chigaji/maron_speak/internal/repository/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer() *echo.Echo {
	cfg := config.LoadConfig()
	e := echo.New()

	//connnect to the db

	db, err := db.Connect(cfg)
	if err != nil {
		e.Logger.Fatal(err)
	}

	//repositories
	userRepo := repo.NewuserRepo(db)
	chatRepo := repo.NewChartRepo(db)
	progressRepo := repo.NewProgressRepo(db)

	//services
	authService := services.NewAuthService(userRepo)
	chatService := services.NewChatService(chatRepo)
	progressSercice := services.NewProgressService(progressRepo)

	//handlers
	authHandler := handlers.NewAuthHandler(authService)
	chatHander := handlers.NewChatHandler(chatService)
	progressHandler := handlers.NewProgressHandler(progressSercice)

	//routes
	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)

	//Authenticated routes

	// jwtMiddleware := middleware.JWTMiddleWare(cfg)

	// e.Use(jwtMiddleware)
	// e.Use(middleware.UserContextMiddleware)
	e.Use(middleware.Logger())

	e.POST("/chat", chatHander.SaveChat)
	e.GET("/chat", chatHander.GetChatHistory)

	e.POST("/progress", progressHandler.UpdateProgress)
	e.GET("/progress", progressHandler.GetProgress)
	return e
}
