package home_page

import (
	"chaincue-real-estate-go/internal/models"
	"chaincue-real-estate-go/internal/services/dto_builder_helpers"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"sort"
	"sync"
)

type HomePageDTO struct {
	Countries           []CountryDTO `json:"countries"`
	RecentlyAddedHouses []HouseDTO   `json:"recentlyAddedHouses"`
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

func RegisterHomePageRoutes(router *gin.Engine) {
	router.GET("/home", homePage)
}

func homePage(c *gin.Context) {
	log.Println("homePage")
	dto := buildDTO(func(builder *DTOBuilder) {})
	c.JSON(200, dto)
}

func buildDTO(additionalProcessing func(*DTOBuilder)) HomePageDTO {
	dtoBuilder := DTOBuilder{}

	if additionalProcessing != nil {
		additionalProcessing(&dtoBuilder)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	countriesChannel := make(chan []models.Country)
	housesChannel := make(chan []models.House)

	go func() {
		defer wg.Done()
		dto_builder_helpers.UpdateDTOBuilderWithCountries(func(dtoBuilder *DTOBuilder, countries []models.Country) {
			dtoBuilder.Countries = countries
		})(&dtoBuilder)
		close(countriesChannel)
	}()

	go func() {
		defer wg.Done()
		dto_builder_helpers.UpdateDTOBuilderWithHouses(func(dtoBuilder *DTOBuilder, houses []models.House) {
			sort.Slice(houses, func(i, j int) bool {
				return houses[i].CreatedAt.After(houses[j].CreatedAt)
			})
			numHouses := int(math.Min(6, float64(len(houses))))
			dtoBuilder.Houses = houses[:numHouses]
		})(&dtoBuilder)
		close(housesChannel)
	}()

	go func() {
		wg.Wait()
	}()
	_, _ = <-countriesChannel, <-housesChannel

	return toHomePageDTO(dtoBuilder)
}

func toHomePageDTO(dtoBuilder DTOBuilder) HomePageDTO {
	return HomePageDTO{
		Countries:           convertCountries(dtoBuilder.Countries),
		RecentlyAddedHouses: convertHouses(dtoBuilder.Houses),
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
