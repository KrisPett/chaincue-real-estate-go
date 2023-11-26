package account_page

import (
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
	router.GET("/account", accountPage)
}

func accountPage(c *gin.Context) {
	log.Println("accountPage")
	toDTO := updateDTOBuilder()
	c.JSON(200, toDTO)
}

func updateDTOBuilder() AccountPageDTO {
	dtoBuilder := DTOBuilder{}

	return toAccountPageDTO(dtoBuilder)
}

func toAccountPageDTO(dtoBuilder DTOBuilder) AccountPageDTO {
	return AccountPageDTO{
		ID: "dtoBuilder.id",
	}
}
