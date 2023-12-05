package houses_page

import (
	"chaincue-real-estate-go/internal/models"
	"chaincue-real-estate-go/internal/services"
	"chaincue-real-estate-go/internal/services/dto_builder_helpers"
	"chaincue-real-estate-go/internal/utilities"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"sync"
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
	Type        string `json:"type"`
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

type FilterSearchReqBody struct {
	Country             string   `json:"country"`
	TextAreaSearchValue string   `json:"textAreaSearchValue"`
	HouseTypes          []string `json:"houseTypes"`
}

func RegisterHousesPageRoutes(router *gin.Engine) {
	router.GET("/houses", housesPage)
	router.PUT("/houses", searchHouses)
}

func housesPage(c *gin.Context) {
	log.Println("housesPage")
	dto := buildDTO(func(builder *DTOBuilder) {})
	c.JSON(200, dto)
}

func searchHouses(c *gin.Context) {
	houseService := services.UseHouseService()
	log.Println("searchHouses")

	body, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{"error": "Error reading request body"})
		return
	}

	var reqBody FilterSearchReqBody
	if err := json.Unmarshal(body, &reqBody); err != nil {
		c.JSON(400, gin.H{"error": "Error decoding JSON data"})
		return
	}

	houses, err := houseService.SearchHouses(reqBody.Country, reqBody.TextAreaSearchValue, reqBody.HouseTypes)
	dtos := convertHouses(houses)
	c.JSON(200, dtos)
}

func buildDTO(additionalProcessing func(*DTOBuilder)) HousesPageDTO {
	dtoBuilder := DTOBuilder{}

	if additionalProcessing != nil {
		additionalProcessing(&dtoBuilder)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		dto_builder_helpers.UpdateDTOBuilderWithCountries(func(dtoBuilder *DTOBuilder, countries []models.Country) {
			dtoBuilder.Countries = countries
		})(&dtoBuilder)
	}()

	go func() {
		defer wg.Done()
		dto_builder_helpers.UpdateDTOBuilderWithHouses(func(dtoBuilder *DTOBuilder, houses []models.House) {
			dtoBuilder.Houses = houses
		})(&dtoBuilder)
	}()

	wg.Wait()

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
		Type:        utilities.FormatTitleCaseString(string(house.HouseTypes)),
		NumberRooms: house.NumberRooms,
		Beds:        house.Beds,
		DollarPrice: house.Price,
		CryptoPrice: "₿32.346",
		Src:         house.Src,
	}
}
