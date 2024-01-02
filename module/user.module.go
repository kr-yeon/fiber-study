package module

import (
	"awesomeProject/controller"
	"awesomeProject/service"
	"github.com/gofiber/fiber/v2"
)

type UserModule struct {
	controller controller.UserController
	service    service.UserService
}

func User(app *fiber.App, userService service.UserService) UserModule {
	return UserModule{
		controller: controller.User(app, userService),
		service:    userService,
	}
}
