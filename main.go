package main

import (
	"github.com/czewski/crud-in-go/pkg/handlers"
	"github.com/gofiber/fiber/v2"
)

const port = 8000

func main() {
	app := fiber.New()

	app.Get("/player/:id?", handlers.ReadPlayers)
	app.Post("/player", handlers.CreatePlayers)
	app.Put("/player/:id", handlers.UpdatePlayers)
	app.Delete("/player/:id", handlers.DeletePlayers)
	app.Listen(":8000")
}
