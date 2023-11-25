package dto_builder_helpers

import (
	"chaincue-real-estate-go/internal/models"
	"chaincue-real-estate-go/internal/services"
)

func UpdateDTOBuilderWithCountries[B any](setCountries func(*B, []models.Country)) func(*B) {
	return func(dtoBuilder *B) {
		countryService := services.UseCountryService()
		countries, _ := countryService.FindAll()
		setCountries(dtoBuilder, countries)
	}
}
