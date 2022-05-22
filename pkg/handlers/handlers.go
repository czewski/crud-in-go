package handlers

import "github.com/gofiber/fiber/v2"

func ReadPlayers(c *fiber.Ctx) error {
	return c.SendString("All Bookmarks")
}

func CreatePlayers(c *fiber.Ctx) error { return c.SendString("All Bookmarks") }

func UpdatePlayers(c *fiber.Ctx) error {
	return c.SendString("All Bookmarks")
}

func DeletePlayers(c *fiber.Ctx) error {
	return c.SendString("All Bookmarks")
}
