package configs

import (
	"chaincue-real-estate-go/internal/models"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
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
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDBName := os.Getenv("POSTGRES_DB")
	postgresPort := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		postgresHost, postgresUser, postgresPassword, postgresDBName, postgresPort,
	)

	var err error
	postgresDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
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
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic("Error connecting to Redis")
	}
	fmt.Println("Connected to Redis successfully.")
}
