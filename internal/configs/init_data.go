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
	createHouse(db, broker, utilities.URLFrontImage1, models.VILLA, "Sweden, Stockholm", "SWEDEN", "Stockholm", 969381)
	createHouse(db, broker, utilities.URLFrontImage2, models.VILLA, "Spain, Málaga", "SPAIN", "Málaga", 969382)
	createHouse(db, broker, utilities.URLFrontImage3, models.VILLA, "Sweden, Stockholm", "SWEDEN", "Stockholm", 969383)
	createHouse(db, broker, utilities.URLFrontImage4, models.VILLA, "Spain, Málaga", "SPAIN", "Málaga", 969384)
	createHouse(db, broker, utilities.URLFrontImage5, models.VACATION_HOME, "Spain, Málaga", "SPAIN", "Málaga", 969385)
	createHouse(db, broker, utilities.URLFrontImage6, models.VACATION_HOME, "Sweden, Stockholm", "SWEDEN", "Stockholm", 969386)
	createHouse(db, broker, utilities.URLFrontImage1, models.VACATION_HOME, "Spain, Málaga", "SPAIN", "Málaga", 969387)
	createHouse(db, broker, utilities.URLFrontImage2, models.ESTATES_AND_FARMS, "Spain, Málaga", "SPAIN", "Málaga", 969388)
	createHouse(db, broker, utilities.URLFrontImage3, models.ESTATES_AND_FARMS, "Sweden, Stockholm", "SWEDEN", "Stockholm", 969389)
	createHouse(db, broker, utilities.URLFrontImage4, models.ESTATES_AND_FARMS, "Sweden, Stockholm", "SWEDEN", "Stockholm", 969314)
	createHouse(db, broker, utilities.URLFrontImage5, models.LAND, "Spain, Málaga", "SPAIN", "Málaga", 969324)
	createHouse(db, broker, utilities.URLFrontImage6, models.LAND, "Sweden, Stockholm", "SPAIN", "Málaga", 969334)
	createHouse(db, broker, utilities.URLFrontImage1, models.LAND, "Spain, Málaga", "SPAIN", "Málaga", 969334)
	createHouse(db, broker, utilities.URLFrontImage2, models.OTHER_HOUSES, "Spain, Málaga", "SPAIN", "Málaga", 969344)
	createHouse(db, broker, utilities.URLFrontImage3, models.OTHER_HOUSES, "Sweden, Stockholm", "SWEDEN", "Stockholm", 969354)
	createHouse(db, broker, utilities.URLFrontImage4, models.OTHER_HOUSES, "Spain, Málaga", "SPAIN", "Málaga", 969364)
	createHouse(db, broker, utilities.URLFrontImage5, models.TOWNHOUSE, "Spain, Málaga", "SPAIN", "Málaga", 969374)
	createHouse(db, broker, utilities.URLFrontImage6, models.TOWNHOUSE, "Sweden, Stockholm", "SWEDEN", "Stockholm", 969384)
	createHouse(db, broker, utilities.URLFrontImage1, models.TOWNHOUSE, "Spain, Málaga", "SPAIN", "Málaga", 969394)
	createHouse(db, broker, utilities.URLFrontImage2, models.CONDOMINIUM, "Spain, Málaga", "SPAIN", "Málaga", 969184)
	createHouse(db, broker, utilities.URLFrontImage3, models.CONDOMINIUM, "Sweden, Stockholm", "SWEDEN", "Stockholm", 969284)
	createHouse(db, broker, utilities.URLFrontImage4, models.CONDOMINIUM, "Spain, Málaga", "SPAIN", "Málaga", 969384)
	createHouse(db, broker, utilities.URLFrontImage5, models.CONDOMINIUM, "Spain, Málaga", "SPAIN", "Málaga", 969484)
}

func createHouse(db *gorm.DB, broker *models.Broker, url string, houseType models.HouseTypes, location string, country string, city string, price int) {
	description := "Welcome to this bright and well-planned four-bedroom apartment with a balcony in a private location and a view of greenery! The residence features well-organized rooms and substantial windows in three different directions, providing a delightful infusion of natural light throughout the entire apartment. You'll find a spacious living room with comfortable seating areas and access to the pleasant balcony, offering sunny exposure and a lovely view of the green surroundings. Additionally, the apartment boasts a spacious kitchen with room for a dining area for the whole family, and here too, you can enjoy a pleasant view of the green area outside.\n\nThis well-planned apartment includes three good-sized bedrooms. Conveniently, for larger families, it offers both a fully tiled bathroom with a washing machine and a guest WC. Ample storage options are available through closets and a walk-in closet.\n\nYou are warmly welcome to visit!"
	house := models.NewHouse(houseType, url)
	house.Location = location
	house.Description = description
	house.Country = country
	house.City = city
	house.NumberRooms = 3
	house.Beds = 2
	house.Price = price
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
