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

func InitAuthModule(app *fiber.App, authService service.AuthService) AuthModule {
	return AuthModule{
		controller: controller.InitAuthController(app, authService),
		service:    authService,
	}
}
