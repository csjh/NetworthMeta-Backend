package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func Connect() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Couldn't read .env file")
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic("⛒ Connection Failed to Database")
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic("⛒ Connection Failed to Database")
	}
	fmt.Println("⛁ Connected to Database")
	return client
}

func Disconnect(client *mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	client = nil

	fmt.Println("⛁ Disconnected from Database")
}