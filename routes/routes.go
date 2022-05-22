package routes

import (
	"github.com/czewski/crud-in-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouteManager(app *fiber.App) {

	app.Get("/player/:id?", controllers.ReadPlayers)
	app.Post("/player", controllers.CreatePlayers)
	app.Put("/player/:id", controllers.UpdatePlayers)
	app.Delete("/player/:id", controllers.DeletePlayers)

}
