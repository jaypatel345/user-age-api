package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaypatel/user-age-api/internal/handler"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users", userHandler.ListUsers)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Delete("/users/:id", userHandler.DeleteUser)
}
