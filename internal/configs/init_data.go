package configs

import (
	"chaincue-real-estate-go/internal/models"
	"chaincue-real-estate-go/internal/utilities"
	"gorm.io/gorm"
	"log"
)

func InitData() {
	db := GetPostgresDB()
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
	createHouse(db, broker, utilities.URLFrontImage1, models.VILLA, "Sweden, Stockholm", "SWEDEN", "Stockholm")
	createHouse(db, broker, utilities.URLFrontImage2, models.VILLA, "Spain, Málaga", "SPAIN", "Málaga")
	createHouse(db, broker, utilities.URLFrontImage3, models.VILLA, "Sweden, Stockholm", "SWEDEN", "Stockholm")
	createHouse(db, broker, utilities.URLFrontImage4, models.VILLA, "Spain, Málaga", "SPAIN", "Málaga")
	createHouse(db, broker, utilities.URLFrontImage5, models.VACATION_HOME, "Spain, Málaga", "SPAIN", "Málaga")
	createHouse(db, broker, utilities.URLFrontImage6, models.VACATION_HOME, "Sweden, Stockholm", "SWEDEN", "Stockholm")
	createHouse(db, broker, utilities.URLFrontImage1, models.VACATION_HOME, "Spain, Málaga", "SPAIN", "Málaga")
	createHouse(db, broker, utilities.URLFrontImage2, models.ESTATES_AND_FARMS, "Spain, Málaga", "SPAIN", "Málaga")
	createHouse(db, broker, utilities.URLFrontImage3, models.ESTATES_AND_FARMS, "Sweden, Stockholm", "SWEDEN", "Stockholm")
	createHouse(db, broker, utilities.URLFrontImage4, models.ESTATES_AND_FARMS, "Sweden, Stockholm", "SWEDEN", "Stockholm")
	createHouse(db, broker, utilities.URLFrontImage5, models.LAND, "Spain, Málaga", "SPAIN", "Málaga")
	createHouse(db, broker, utilities.URLFrontImage6, models.LAND, "Sweden, Stockholm", "SPAIN", "Málaga")
	createHouse(db, broker, utilities.URLFrontImage1, models.LAND, "Spain, Málaga", "SPAIN", "Málaga")
	createHouse(db, broker, utilities.URLFrontImage2, models.OTHER_HOUSES, "Spain, Málaga", "SPAIN", "Málaga")
	createHouse(db, broker, utilities.URLFrontImage3, models.OTHER_HOUSES, "Sweden, Stockholm", "SWEDEN", "Stockholm")
	createHouse(db, broker, utilities.URLFrontImage4, models.OTHER_HOUSES, "Spain, Málaga", "SPAIN", "Málaga")
	createHouse(db, broker, utilities.URLFrontImage5, models.TOWNHOUSE, "Spain, Málaga", "SPAIN", "Málaga")
	createHouse(db, broker, utilities.URLFrontImage6, models.TOWNHOUSE, "Sweden, Stockholm", "SWEDEN", "Stockholm")
	createHouse(db, broker, utilities.URLFrontImage1, models.TOWNHOUSE, "Spain, Málaga", "SPAIN", "Málaga")
	createHouse(db, broker, utilities.URLFrontImage2, models.CONDOMINIUM, "Spain, Málaga", "SPAIN", "Málaga")
	createHouse(db, broker, utilities.URLFrontImage3, models.CONDOMINIUM, "Sweden, Stockholm", "SWEDEN", "Stockholm")
	createHouse(db, broker, utilities.URLFrontImage4, models.CONDOMINIUM, "Spain, Málaga", "SPAIN", "Málaga")
	createHouse(db, broker, utilities.URLFrontImage5, models.CONDOMINIUM, "Spain, Málaga", "SPAIN", "Málaga")
}

func createHouse(db *gorm.DB, broker *models.Broker, url string, houseType models.HouseTypes, location string, country string, city string) {
	house := models.NewHouse(houseType, url)
	house.Location = location
	house.Country = country
	house.City = city
	house.NumberRooms = 3
	house.Beds = 2
	house.Price = "$969 384"
	saveHouseWithImages(db, house, broker)
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
