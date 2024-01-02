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

func InitUserModule(app *fiber.App, userService service.UserService) UserModule {
	return UserModule{
		controller: controller.InitUserController(app, userService),
		service:    userService,
	}
}
