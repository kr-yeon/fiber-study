package module

import (
	"awesomeProject/service"
	"github.com/gofiber/fiber/v2"
)

func InitAppModule(app *fiber.App) {
	userService := service.UserService{}
	authService := service.AuthService{}

	InitUserModule(app, userService)
	InitAuthModule(app, authService)
}
