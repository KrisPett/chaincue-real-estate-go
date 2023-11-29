package configs

import (
	"chaincue-real-estate-go/internal/models"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresDB *gorm.DB

var redisClient *redis.Client

func GetPostgresDB() *gorm.DB {
	return postgresDB
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func ConnectPostgres() {
	var err error
	postgresDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=admin password=admin dbname=chaincue-real-estate-db port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = postgresDB.AutoMigrate(
		&models.House{},
		&models.HouseImage{},
		&models.Broker{},
		&models.Country{},
	)

	if err != nil {
		panic("failed to migrate database")
	}
	fmt.Println("Database migration completed successfully.")
}

func ConnectRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic("Error connecting to Redis")
	}
	fmt.Println("Connected to Redis successfully.")
}
