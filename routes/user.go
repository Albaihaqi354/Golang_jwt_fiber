package routes

import (
	"golang_jwt_copy/controllers"
	"golang_jwt_copy/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app fiber.Router) {
	users := app.Group("/users")

	users.Use(middleware.Auth)

	users.Get("/me", controllers.Me)
}
