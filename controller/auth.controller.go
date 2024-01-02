package controller

import (
	"awesomeProject/service"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	router  fiber.Router
	service service.AuthService
}

func Auth(app *fiber.App, authService service.AuthService) AuthController {
	controller := AuthController{
		router:  app.Group("auth"),
		service: authService,
	}

	controller.getAuth()

	return controller
}

// @tags Auth
// @router /auth [get]
func (controller AuthController) getAuth() {
	controller.router.Get("/", controller.service.GetAuth)
}
