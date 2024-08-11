package controllers

import (
	"golang_jwt_copy/helpers"
	"golang_jwt_copy/models"

	"github.com/gofiber/fiber/v2"
)

func Me(c *fiber.Ctx) error {
	user := c.Locals("userinfo").(*helpers.MyCutomClaims)

	userResponse := &models.MyProfile{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "My Profile",
		"profile": userResponse,
	})
}
