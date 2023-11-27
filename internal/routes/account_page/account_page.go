package account_page

import (
	"chaincue-real-estate-go/internal/configs"
	"github.com/gin-gonic/gin"
	"log"
)

type AccountPageDTO struct {
	ID string `json:"id"`
}

type DTOBuilder struct {
	id string
}

func RegisterAccountPageRoutes(router *gin.Engine) {
	router.GET("/account", configs.ProtectRoute, accountPage)
	/*TODO */
	//router.Use(configs.ProtectRoute)
}

func accountPage(c *gin.Context) {
	log.Println("accountPage")
	dto := buildDTO(func(builder *DTOBuilder) {})
	c.JSON(200, dto)
}

func buildDTO(additionalProcessing func(*DTOBuilder)) AccountPageDTO {
	dtoBuilder := DTOBuilder{}

	if additionalProcessing != nil {
		additionalProcessing(&dtoBuilder)
	}

	return toAccountPageDTO(dtoBuilder)
}

func toAccountPageDTO(dtoBuilder DTOBuilder) AccountPageDTO {
	return AccountPageDTO{
		ID: "dtoBuilder.id",
	}
}
