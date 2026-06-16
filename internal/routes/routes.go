package routes

import (
	"github.com/gofiber/fiber/v2"

	"ainyx/internal/handler"
)

func RegisterRoutes(
	app *fiber.App,
	userHandler *handler.UserHandler,
) {
	app.Post("/users", userHandler.CreateUser)

	app.Get("/users/:id", userHandler.GetUser)

	app.Get("/users", userHandler.GetAllUsers)

	app.Put("/users/:id", userHandler.UpdateUser)

	app.Delete("/users/:id", userHandler.DeleteUser)
}
