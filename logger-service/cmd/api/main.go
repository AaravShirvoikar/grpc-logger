package main

import (
	"context"
	"log"
	"time"

	"github.com/AaravShirvoikar/logger-service/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const port = "50001"
const url = "mongodb://mongo:27017"

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	mongoClient, err := connect()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	app.gRPCListen()
}

func connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(url)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("error connecting:", err)
		return nil, err
	}

	log.Println("connected to mongo")

	return c, nil
}
