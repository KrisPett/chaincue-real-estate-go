package configs

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func SetupCORS(router *gin.Engine) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	corsAllowOrigins := os.Getenv("CORS_ALLOW_ORIGINS")
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{corsAllowOrigins}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "X-CSRF-Token"}

	router.Use(cors.New(config))
}
