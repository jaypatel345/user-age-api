package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	db "github.com/jaypatel/user-age-api/db/sqlc"
	"github.com/jaypatel/user-age-api/internal/handler"
	"github.com/jaypatel/user-age-api/internal/repository"
	"github.com/jaypatel/user-age-api/internal/routes"
	"github.com/jaypatel/user-age-api/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	database, err := sql.Open("postgres", databaseURL)

	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	if err := database.Ping(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	queries := db.New(database)
	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	routes.SetupRoutes(app, userHandler)
	log.Fatal(app.Listen(":8080"))
}
