package dto_builder_helpers

import (
	"chaincue-real-estate-go/internal/models"
	"chaincue-real-estate-go/internal/services"
)

func UpdateDTOBuilderWithHouses[B any](setHouses func(*B, []models.House)) func(*B) {
	return func(dtoBuilder *B) {
		houseService := services.UseHouseService()
		houses, _ := houseService.FindAll()
		setHouses(dtoBuilder, houses)
	}
}

func UpdateDTOBuilderWithHouseByHouseId[B any](houseId string, setHouse func(*B, models.House)) func(*B) {
	return func(dtoBuilder *B) {
		houseService := services.UseHouseService()
		house, _ := houseService.FindById(houseId)
		setHouse(dtoBuilder, house)
	}
}
