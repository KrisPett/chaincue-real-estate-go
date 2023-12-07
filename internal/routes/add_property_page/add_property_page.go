package add_property_page

import (
	"chaincue-real-estate-go/internal/models"
	"chaincue-real-estate-go/internal/services"
	"chaincue-real-estate-go/internal/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CreatePropertyReqBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Supply      string `json:"supply"`
}

func RegisterAddPropertyPageRoutes(router *gin.RouterGroup) {
	router.POST("/add-property", addProperty)
}

func addProperty(c *gin.Context) {
	log.Println("Adding a new property")

	var reqBody CreatePropertyReqBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		log.Printf("Error decoding JSON data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode the request body"})
		return
	}

	authHeader := c.GetHeader("Authorization")
	token := utilities.TrimAndGetToken(authHeader)
	fmt.Print(token)

	houseService := services.UseHouseService()
	house := models.NewHouse("CONDOMINIUM", utilities.URLFrontImage1)
	house.Title = reqBody.Title
	house.Description = reqBody.Description

	if err := houseService.Create(house); err != nil {
		log.Printf("Failed to create a house: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "House creation failed"})
		return
	}

	imageURLs := []string{
		house.Src,
		utilities.URL1,
		utilities.URL2,
		utilities.URL3,
		utilities.URL4,
		utilities.URL5,
		utilities.URL6,
	}
	houseImageService := services.UseHouseImageService()
	for _, url := range imageURLs {
		houseImage := models.NewHouseImage(url, house.ID)
		err := houseImageService.Create(houseImage)
		if err != nil {
			log.Printf("Failed to create a house image: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "House image creation failed"})
			return
		}
	}

	c.JSON(http.StatusOK, house.ID)
}
