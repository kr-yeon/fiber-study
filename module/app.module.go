package module

import (
	"awesomeProject/service"
	"github.com/gofiber/fiber/v2"
)

func App(app *fiber.App) {
	userService := service.UserService{}
	authService := service.AuthService{}

	User(app, userService)
	Auth(app, authService)
}
