package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/czewski/crud-in-go/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReadPlayers(c *fiber.Ctx) error {
	client := models.ConnectDB()
	collection := client.Database("filas").Collection("crud")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("userId")

	player := models.Player{}
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	err := collection.FindOne(ctx, bson.M{"id": objId}).Decode(&player)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(models.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": player}})

}

func CreatePlayers(c *fiber.Ctx) error {
	client := models.ConnectDB()
	collection := client.Database("filas").Collection("crud")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	player := models.Player{}
	if err := c.BodyParser(&player); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	newPlayer := models.Player{
		Name: player.Name,
		Nick: player.Nick,
		Team: player.Team,
		Game: player.Game,
		Age:  player.Age,
	}

	result, err := collection.InsertOne(ctx, newPlayer)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(models.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func UpdatePlayers(c *fiber.Ctx) error {
	client := models.ConnectDB()
	collection := client.Database("filas").Collection("crud")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	userId := c.Params("userId")
	player := models.Player{}
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	if err := c.BodyParser(&player); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	update := bson.M{"name": player.Name, "nick": player.Nick, "team": player.Team, "game": player.Game, "age": player.Age}
	result, err := collection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	var updatedUser models.Player
	if result.MatchedCount == 1 {
		err := collection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(models.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(models.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedUser}})
}

func DeletePlayers(c *fiber.Ctx) error {
	client := models.ConnectDB()
	collection := client.Database("filas").Collection("crud")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("userId")
	defer cancel()

	//objId, _ := primitive.ObjectIDFromHex(userId)

	result, err := collection.DeleteOne(ctx, bson.M{"name": userId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			models.UserResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "Player not found!"}},
		)
	}
	return c.Status(http.StatusOK).JSON(
		models.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Player deleted!"}},
	)
}
