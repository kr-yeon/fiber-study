package main

import (
	_ "awesomeProject/docs"
	"awesomeProject/module"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"os"
)

// @title Fiber Example API
// @version 1.0
// @BasePath /
func main() {
	app := fiber.New()

	if os.Getenv("ENABLE_SWAGGER") == "1" {
		app.Get("/docs/*", swagger.HandlerDefault)
	}

	module.App(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
