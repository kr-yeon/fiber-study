package controller

import (
	"awesomeProject/service"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	router  fiber.Router
	service service.UserService
}

func User(app *fiber.App, userService service.UserService) UserController {
	controller := UserController{
		router:  app.Group("user"),
		service: userService,
	}

	controller.getUser()

	return controller
}

// @tags User
// @router /user [get]
func (controller UserController) getUser() {
	controller.router.Get("/", controller.service.GetUser)
}
