package routes

import (
	"golang_jwt_copy/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router) {
	auth := app.Group("/auth")

	auth.Post("/register", controllers.Register)

	auth.Post("/login", controllers.Login)
}