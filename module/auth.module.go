package module

import (
	"awesomeProject/controller"
	"awesomeProject/service"
	"github.com/gofiber/fiber/v2"
)

type AuthModule struct {
	controller controller.AuthController
	service    service.AuthService
}

func Auth(app *fiber.App, authService service.AuthService) AuthModule {
	return AuthModule{
		controller: controller.Auth(app, authService),
		service:    authService,
	}
}
