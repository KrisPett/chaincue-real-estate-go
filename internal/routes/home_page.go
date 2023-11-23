package routes

import (
	"chaincue-real-estate-go/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterHomePageRoutes(router *gin.Engine) {
	router.GET("/home", homePage)
}

func homePage(c *gin.Context) {
	//brokerService := services.UseBrokerService()
	countryService := services.UseCountryService()
	houseImageService := services.UseHouseImageService()
	houseService := services.UseHouseService()

	findAll, _ := houseService.FindAll()

	_, err := houseImageService.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_, err2 := countryService.FindAll()
	if err2 != nil {
		c.JSON(500, gin.H{"error": err2.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Hello World!", "countries": findAll})
}
