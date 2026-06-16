package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaypatel/user-age-api/internal/handler"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	user := app.Group("/users")
	user.Post("/", userHandler.CreateUser)
	user.Get("/", userHandler.ListUsers)
	user.Get("/:id", userHandler.GetUserByID)
	user.Put("/:id", userHandler.UpdateUser)
	user.Delete("/:id", userHandler.DeleteUser)
}
