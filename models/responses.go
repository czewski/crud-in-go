package models

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

type Player struct {
	Id   primitive.ObjectID `json:"id,omitempty"`
	Name string             `json:"name,omitempty"`
	Nick string             `json:"nick,omitempty"`
	Team string             `json:"team,omitempty"`
	Game string             `json:"game,omitempty"`
	Age  int                `json:"age,omitempty"`
}
