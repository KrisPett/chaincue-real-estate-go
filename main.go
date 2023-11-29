package main

import (
	"chaincue-real-estate-go/internal/configs"
	"chaincue-real-estate-go/internal/routes/account_page"
	"chaincue-real-estate-go/internal/routes/home_page"
	"chaincue-real-estate-go/internal/routes/house_page"
	"chaincue-real-estate-go/internal/routes/houses_page"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectPostgres()
	configs.ConnectRedis()
	configs.InitData()

	router := gin.Default()
	configs.SetupCORS(router)

	//AnonymousRoutes
	home_page.RegisterHomePageRoutes(router)
	houses_page.RegisterHousesPageRoutes(router)
	house_page.RegisterHousePageRoutes(router)

	//UserAuthenticatedRoutes
	userRoutes := router.Group("/user")
	userRoutes.Use(func(c *gin.Context) {
		configs.AuthenticateRoutes(c, "user")
	})
	{
		account_page.RegisterAccountPageRoutes(userRoutes)
	}

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error:", err)
	}
}
