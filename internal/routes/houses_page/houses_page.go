package houses_page

import (
	"chaincue-real-estate-go/internal/models"
	"chaincue-real-estate-go/internal/services/dto_builder_helpers"
	"github.com/gin-gonic/gin"
	"log"
)

type HousesPageDTO struct {
	Countries []CountryDTO `json:"countries"`
	Houses    []HouseDTO   `json:"houses"`
}

type CountryDTO struct {
	Name string `json:"name"`
}

type HouseDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Location    string `json:"location"`
	NumberRooms int    `json:"numberRooms"`
	Beds        int    `json:"beds"`
	DollarPrice string `json:"dollarPrice"`
	CryptoPrice string `json:"cryptoPrice"`
	Src         string `json:"src"`
}

type DTOBuilder struct {
	Countries []models.Country
	Houses    []models.House
}

func RegisterHousesPageRoutes(router *gin.Engine) {
	router.GET("/houses", housesPage)
}

func housesPage(c *gin.Context) {
	log.Println("housesPage")
	dto := buildDTO(nil)
	c.JSON(200, dto)
}

func buildDTO(additionalProcessing func(*DTOBuilder)) HousesPageDTO {
	dtoBuilder := DTOBuilder{}

	if additionalProcessing != nil {
		additionalProcessing(&dtoBuilder)
	}

	dto_builder_helpers.UpdateDTOBuilderWithCountries(func(dtoBuilder *DTOBuilder, countries []models.Country) {
		dtoBuilder.Countries = countries
	})(&dtoBuilder)

	dto_builder_helpers.UpdateDTOBuilderWithHouses(func(dtoBuilder *DTOBuilder, houses []models.House) {
		dtoBuilder.Houses = houses
	})(&dtoBuilder)

	return toHomePageDTO(dtoBuilder)
}

func toHomePageDTO(dtoBuilder DTOBuilder) HousesPageDTO {
	return HousesPageDTO{
		Countries: convertCountries(dtoBuilder.Countries),
		Houses:    convertHouses(dtoBuilder.Houses),
	}
}

func convertCountries(countries []models.Country) []CountryDTO {
	var result []CountryDTO
	for _, country := range countries {
		result = append(result, toCountryDTO(country))
	}
	return result
}

func convertHouses(houses []models.House) []HouseDTO {
	var result []HouseDTO
	for _, house := range houses {
		result = append(result, toHouseDTO(house))
	}
	return result
}

func toCountryDTO(country models.Country) CountryDTO {
	return CountryDTO{
		Name: string(country.Name),
	}
}

func toHouseDTO(house models.House) HouseDTO {
	return HouseDTO{
		ID:          house.ID,
		Title:       house.Title,
		Location:    house.Location,
		NumberRooms: house.NumberRooms,
		Beds:        house.Beds,
		DollarPrice: house.Price,
		CryptoPrice: "₿32.346",
		Src:         house.Src,
	}
}
