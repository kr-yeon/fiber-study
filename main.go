package main

import (
	_ "awesomeProject/docs"
	"awesomeProject/module"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Fiber Example API
// @version 1.0
// @BasePath /
func main() {
	app := fiber.New()
	app.Get("/docs/*", swagger.HandlerDefault)

	module.InitAppModule(app)

	app.Listen(":3000")
}
