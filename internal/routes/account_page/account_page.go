package account_page

import (
	"chaincue-real-estate-go/internal/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type AccountPageDTO struct {
	ID string `json:"id"`
}

type DTOBuilder struct {
	id string
}

func RegisterAccountPageRoutes(router *gin.RouterGroup) {
	router.GET("/account", accountPage)
}

func accountPage(c *gin.Context) {
	log.Println("accountPage")
	authHeader := c.GetHeader("Authorization")
	token := utilities.TrimAndGetToken(authHeader)
	dto := buildDTO(token, func(builder *DTOBuilder) {})
	c.JSON(200, dto)
}

func buildDTO(token string, additionalProcessing func(*DTOBuilder)) AccountPageDTO {
	fmt.Println(token)
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
