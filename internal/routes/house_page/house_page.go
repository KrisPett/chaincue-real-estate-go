package house_page

import (
	"chaincue-real-estate-go/internal/models"
	"chaincue-real-estate-go/internal/services/dto_builder_helpers"
	"chaincue-real-estate-go/internal/utilities"
	"github.com/gin-gonic/gin"
	"log"
)

type HousePageDTO struct {
	ID            string          `json:"id"`
	Title         string          `json:"title"`
	Type          string          `json:"type"`
	Location      string          `json:"location"`
	NumberOfRooms int             `json:"numberOfRooms"`
	Beds          int             `json:"beds"`
	DollarPrice   string          `json:"dollarPrice"`
	CryptoPrice   string          `json:"cryptoPrice"`
	Description   string          `json:"description"`
	Images        []HouseImageDTO `json:"images"`
	Broker        *BrokerDTO      `json:"broker"`
}

type HouseImageDTO struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type BrokerDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}

type DTOBuilder struct {
	House models.House
}

func RegisterHousePageRoutes(router *gin.Engine) {
	router.GET("/house/:houseId", housePage)
}

func housePage(c *gin.Context) {
	log.Println("housePage")
	houseId := c.Param("houseId")
	dto := buildDTO(nil, houseId)
	c.JSON(200, dto)
}

func buildDTO(additionalProcessing func(*DTOBuilder), houseId string) HousePageDTO {
	dtoBuilder := DTOBuilder{}

	if additionalProcessing != nil {
		additionalProcessing(&dtoBuilder)
	}

	dto_builder_helpers.UpdateDTOBuilderWithHouseByHouseId(houseId, func(dtoBuilder *DTOBuilder, house models.House) {
		dtoBuilder.House = house
	})(&dtoBuilder)

	return toHomePageDTO(dtoBuilder)
}

func toHomePageDTO(dtoBuilder DTOBuilder) HousePageDTO {
	return HousePageDTO{
		ID:            dtoBuilder.House.ID,
		Title:         dtoBuilder.House.Title,
		Type:          utilities.FormatTitleCaseString(string(dtoBuilder.House.HouseTypes)),
		Location:      dtoBuilder.House.Location,
		NumberOfRooms: dtoBuilder.House.NumberRooms,
		Beds:          dtoBuilder.House.Beds,
		DollarPrice:   dtoBuilder.House.Price,
		CryptoPrice:   "₿32.346",
		Description:   dtoBuilder.House.Description,
		Images:        convertHouseImages(dtoBuilder.House.HouseImages),
		Broker:        toBrokerDTO(dtoBuilder.House.Broker),
	}
}

func convertHouseImages(countries []models.HouseImage) []HouseImageDTO {
	var result []HouseImageDTO
	for _, houseImage := range countries {
		result = append(result, toHouseImageDTO(houseImage))
	}
	return result
}

func toBrokerDTO(broker *models.Broker) *BrokerDTO {
	if broker == nil {
		return nil
	}
	return &BrokerDTO{
		ID:          broker.ID,
		Name:        broker.Name,
		PhoneNumber: broker.PhoneNumber,
		Email:       broker.Email,
	}
}

func toHouseImageDTO(houseImage models.HouseImage) HouseImageDTO {
	return HouseImageDTO{
		ID:  houseImage.ID,
		URL: houseImage.URL,
	}
}
