package main

import (
	"chaincue-real-estate-go/internal/configs"
	"chaincue-real-estate-go/internal/domains"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		db := configs.GetDB()
		broker := domains.NewBroker("d")
		db.Create(broker)
		country := domains.NewCountry(domains.SWEDEN)
		db.Create(country)

		//house := domains.NewHouse(domains.VILLA, utilities.URL1)
		//db.Create(house)
		//
		//houseImage := domains.NewHouseImage(utilities.URL1, house.ID)
		//db.Create(houseImage)

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error:", err)
	}
}
