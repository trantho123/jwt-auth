package main

import (
	"Jwtwithecdsa/api/cmd/routes"
	"Jwtwithecdsa/api/internal/controller"
	"Jwtwithecdsa/api/internal/handler"
	"Jwtwithecdsa/api/internal/repository"
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	println("Start")
	LoadEnv()
	collection, err := ConnectToDatabase()
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	repo := repository.New(collection)
	ctrl := controller.New(repo)
	rtr := routes.New(
		handler.New(ctrl),
	)
	app := fiber.New()
	rtr.Routes(app)
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

func ConnectToDatabase() (*mongo.Collection, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_SERVER")))
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return client.Database(os.Getenv("MONGO_DATABSAE")).Collection(os.Getenv("MONGO_COLLECTION")), nil
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
}
