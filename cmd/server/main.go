package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	db "github.com/jaypatel/user-age-api/db/sqlc"
	"github.com/jaypatel/user-age-api/internal/handler"
	"github.com/jaypatel/user-age-api/internal/logger"
	"github.com/jaypatel/user-age-api/internal/repository"
	"github.com/jaypatel/user-age-api/internal/routes"
	"github.com/jaypatel/user-age-api/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger()
	defer logger.Sync()

	if err := godotenv.Load(); err != nil {
		logger.Logger.Warn("No .env file found", zap.Error(err))
	}

	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		logger.Logger.Fatal("DATABASE_URL is required")
	}

	database, err := sql.Open("postgres", databaseURL)

	if err != nil {
		logger.Logger.Fatal("Failed to open database connection", zap.Error(err))
	}

	defer database.Close()

	if err := database.Ping(); err != nil {
		logger.Logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	logger.Logger.Info("Database connected")

	app := fiber.New()
	queries := db.New(database)
	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	routes.SetupRoutes(app, userHandler)
	logger.Logger.Info("Server started", zap.String("addr", ":8080"))
	log.Fatal(app.Listen(":8080"))
}
