package controllers

import (
	"golang_jwt_copy/configs"
	"golang_jwt_copy/helpers"
	"golang_jwt_copy/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var register models.Register

	if err := c.BodyParser(&register); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if register.Password != register.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Password does not match",
		})
	}

	passwordHash, err := helpers.HashPassword(register.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := models.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: passwordHash,
	}

	if err := configs.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Register Successfully",
	})
}

func Login(c *fiber.Ctx) error {
	var login models.Login

	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var user models.User
	if err := configs.DB.First(&user, "email = ?", login.Email).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if err := helpers.VerifyPassword(user.Password, login.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Password",
		})
	}

	token, err := helpers.CreateToken(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login Successfully",
		"token":   token,
	})
}
