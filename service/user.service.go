package service

import "github.com/gofiber/fiber/v2"

type UserService struct {
}

func (service UserService) GetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})
}
