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
	port := "3000"

	if os.Getenv("ENABLE_SWAGGER") == "1" {
		app.Get("/docs/*", swagger.HandlerDefault)
	}

	module.App(app)

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}
