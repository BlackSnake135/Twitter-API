package main

import (
	"github.com/gofiber/fiber/v2"

	"twitter/app/middleware"
	"twitter/app/routes"
)

func main() {
	app := fiber.New()

	app.Use(middleware.Logger())

	routes.SetupRoutes(app)

	app.Listen(":3001")
}
