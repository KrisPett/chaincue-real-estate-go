package configs

import (
	"chaincue-real-estate-go/internal/models"
	"chaincue-real-estate-go/internal/utilities"
	"gorm.io/gorm"
	"log"
)

func InitData() {
	db := GetDB()
	result := db.Exec(`DELETE FROM brokers; DELETE FROM house_images; DELETE FROM houses; DELETE FROM countries;`)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	/*Broker*/
	broker := models.NewBroker("John Doe")
	db.Create(broker)

	/*Country*/
	country1 := models.NewCountry(models.SWEDEN)
	country2 := models.NewCountry(models.SPAIN)
	db.Create(country1)
	db.Create(country2)

	/*House*/
	for i := 0; i < 18; i++ {
		house := models.NewHouse(models.VILLA, utilities.URLFrontImage1)
		house.Location = "Spain, Málaga"
		house.NumberRooms = 3
		house.Beds = 2
		house.Price = "$969 384"
		saveHouseWithImages(db, house, broker)
	}
}

func saveHouseWithImages(db *gorm.DB, house *models.House, broker *models.Broker) {
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
		houseImage := models.NewHouseImage(url, house.ID)
		db.Create(houseImage)
	}
}
