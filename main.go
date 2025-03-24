package main

import (
	"myapp/config"
	"myapp/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDatabase()

	app := fiber.New()

	routes.AuthRoutes(app)
	routes.UserRoutes(app)

	app.Listen(":8080")
}
