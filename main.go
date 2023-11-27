package main

import (
	"chaincue-real-estate-go/internal/configs"
	"chaincue-real-estate-go/internal/routes/account_page"
	"chaincue-real-estate-go/internal/routes/home_page"
	"chaincue-real-estate-go/internal/routes/house_page"
	"chaincue-real-estate-go/internal/routes/houses_page"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()
	configs.InitData()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "X-CSRF-Token"}
	router.Use(cors.New(config))

	home_page.RegisterHomePageRoutes(router)
	houses_page.RegisterHousesPageRoutes(router)
	house_page.RegisterHousePageRoutes(router)
	account_page.RegisterAccountPageRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error:", err)
	}
}
