package db

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func Connect() *mongo.Client {
	if os.Getenv("MONGODB_URL") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Couldn't read .env file")
		}
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic("⛒ Connection Failed to Database")
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic("⛒ Connection Failed to Database")
	}

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
}
