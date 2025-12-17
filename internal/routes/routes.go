package routes

import (
	"github.com/gofiber/fiber/v2"

	"Go_Backend_Development_Task/internal/handler"
)

func Register(app *fiber.App, h *handler.UserHandler) {

	// Health
	app.Get("/health", handler.HealthCheck)

	// Users
	users := app.Group("/users")
	users.Post("/", h.CreateUser)
	users.Get("/", h.GetUsers)
	users.Get("/:id", h.GetUser)
	users.Put("/:id", h.UpdateUser)
	users.Delete("/:id", h.DeleteUser)
}
