package models

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (client *mongo.Client) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client(). //ReadKey("mongoPwd")
						ApplyURI("mongodb+srv://go:" + os.Getenv("MONGOPWD") + "@cluster0.gs1ad.mongodb.net/myFirstDatabase?retryWrites=true&w=majority").
						SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
	}
	return client
}
