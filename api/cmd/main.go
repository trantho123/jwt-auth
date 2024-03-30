package main

import (
	"Jwtwithecdsa/api/cmd/routes"
	"Jwtwithecdsa/api/internal/controller"
	"Jwtwithecdsa/api/internal/handler"
	"Jwtwithecdsa/api/internal/rds"
	"Jwtwithecdsa/api/internal/repository"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
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
	rdb, err := ConnectRedis()
	if err != nil {
		log.Fatalf("Failed to ping Redis: %v", err)
	}
	repo := repository.New(collection)
	ctrl := controller.New(repo, rds.New(rdb))
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

func ConnectRedis() (*redis.Client, error) {
	// Tạo một đối tượng redis.Client để kết nối đến máy chủ Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "", // Đặt mật khẩu của Redis (nếu cần)
		DB:       0,  // Chọn cơ sở dữ liệu (mặc định là 0)
	})

	// Kiểm tra kết nối đến Redis
	_, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		// Nếu có lỗi, trả về nil và lỗi tương ứng
		return nil, err
	}

	// In ra "PONG" để xác nhận kết nối thành công
	fmt.Println("Kết nối đến Redis thành công")

	// Trả về đối tượng redis.Client và không có lỗi
	return rdb, nil
}
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
}
