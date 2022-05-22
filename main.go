package main

import (
	"github.com/czewski/crud-in-go/routes"
	"github.com/gofiber/fiber/v2"
)

const port = 8000

func main() {
	app := fiber.New()

	routes.RouteManager(app) //add this

	app.Listen(":8000")
}
