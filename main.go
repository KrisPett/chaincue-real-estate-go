package main

import (
	"chaincue-real-estate-go/internal/configs"
	"chaincue-real-estate-go/internal/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()
	configs.InitData()

	r := gin.Default()
	routes.RegisterHomePageRoutes(r)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error:", err)
	}
}
