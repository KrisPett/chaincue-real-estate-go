package main

import (
	"chaincue-real-estate-go/internal/configs"
	"chaincue-real-estate-go/internal/routes/home_page"
	"chaincue-real-estate-go/internal/routes/house_page"
	"chaincue-real-estate-go/internal/routes/houses_page"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()
	configs.InitData()

	r := gin.Default()
	home_page.RegisterHomePageRoutes(r)
	houses_page.RegisterHousesPageRoutes(r)
	house_page.RegisterHousePageRoutes(r)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error:", err)
	}
}
