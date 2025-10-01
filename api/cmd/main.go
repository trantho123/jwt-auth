package main

import (
	"Jwtwithecdsa/api/cmd/routes"
	"Jwtwithecdsa/api/internal/controller"
	"Jwtwithecdsa/api/internal/handler"
	"Jwtwithecdsa/api/internal/rds"
	"Jwtwithecdsa/api/internal/repository"
	"Jwtwithecdsa/api/internal/utils"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	println("Start")
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Faild to load env file: ", err)
	}
	collection, err := ConnectToDatabase(config)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	rdb, err := ConnectRedis(config)
	if err != nil {
		log.Fatalf("Failed to ping Redis: %v", err)
	}
	repo := repository.New(collection)
	ctrl := controller.New(repo, rds.New(rdb), config)
	rtr := routes.New(
		handler.New(ctrl),
	)
	app := fiber.New()
	rtr.Routes(app)
	log.Fatal(app.Listen(":" + config.PORT))
}

func ConnectToDatabase(config utils.Config) (*mongo.Collection, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MONGO_SERVER))
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return client.Database(config.MONGO_DATABSAE).Collection(config.MONGO_COLLECTION), nil
}

func ConnectRedis(config utils.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "cache:" + config.REDIS_SERVER,
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}
