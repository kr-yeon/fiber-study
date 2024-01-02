package service

import "github.com/gofiber/fiber/v2"

type AuthService struct {
}

func (service AuthService) GetAuth(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})
}
