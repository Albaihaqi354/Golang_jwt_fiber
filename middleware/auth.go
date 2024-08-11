package middleware

import (
	"golang_jwt_copy/helpers"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	accessToken := c.Get("Authorization")

	if accessToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	user, err := helpers.ValidationToken(accessToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	c.Locals("userinfo", user)

	return c.Next()
}
