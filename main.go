package main

import (
	"golang_jwt_copy/configs"
	"golang_jwt_copy/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.ConnectDb()

	app := fiber.New()

	api := app.Group("/api")

	routes.AuthRoutes(api)
	routes.UserRoutes(api)

	log.Println("Server Running on Port: 8080")
	log.Fatal(app.Listen(":8080"))
}
