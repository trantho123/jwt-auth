package repository

import (
	"Jwtwithecdsa/api/internal/utils"
	"context"
	"log"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var repo Registry

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Faild to load env file: ", err)
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MONGO_SERVER))
	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}

	coll := client.Database(config.MONGO_DATABSAE).Collection(config.MONGO_COLLECTION)
	repo = New(coll)
	os.Exit(m.Run())
}
