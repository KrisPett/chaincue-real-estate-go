package main

import (
	"chaincue-real-estate-go/internal/configs"
	"chaincue-real-estate-go/internal/domains"
	"chaincue-real-estate-go/internal/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func main() {
	configs.ConnectDB()
	initData()

	r := gin.Default()
	r.GET("/all", func(c *gin.Context) {
		db := configs.GetDB()
		var brokers []domains.Broker
		result := db.Find(&brokers)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, brokers)
	})

	r.GET("/houses", func(c *gin.Context) {
		db := configs.GetDB()
		var houses []domains.House
		result := db.Preload("HouseImages").Preload("Broker").Find(&houses)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, houses)
	})

	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error:", err)
	}

}

func initData() {
	db := configs.GetDB()
	result := db.Exec(`DELETE FROM brokers; DELETE FROM house_images; DELETE FROM houses; DELETE FROM countries;`)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	/*Broker*/
	broker := domains.NewBroker("d")
	db.Create(broker)

	/*Country*/
	country1 := domains.NewCountry(domains.SWEDEN)
	country2 := domains.NewCountry(domains.SPAIN)
	db.Create(country1)
	db.Create(country2)

	/*House*/
	for i := 0; i < 18; i++ {
		house := domains.NewHouse(domains.VILLA, utilities.URLFrontImage1)
		saveHouseWithImages(db, house, broker)
	}
}

func saveHouseWithImages(db *gorm.DB, house *domains.House, broker *domains.Broker) {
	house.Broker = broker
	db.Create(house)

	imageURLs := []string{
		house.Src,
		utilities.URL1,
		utilities.URL2,
		utilities.URL3,
		utilities.URL4,
		utilities.URL5,
		utilities.URL6,
	}

	for _, url := range imageURLs {
		houseImage := domains.NewHouseImage(url, house.ID)
		db.Create(houseImage)
	}
}
