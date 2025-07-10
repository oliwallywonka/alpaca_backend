package dtos

import (
	"github.com/oliwallywonka/alpaca_backend/db/model"
)

type HotelRoomDTO struct {
	HotelID  string  `json:"hotel_id"`
	Type     string  `json:"type"`
	RefPrice float64 `json:"ref_price"`
	//Capacity  int `json:"capacity"`
}

func HotelRoomModelToDTO(model *model.HotelRoom) *HotelRoomDTO {
	return &HotelRoomDTO{
		HotelID:  model.HotelID,
		Type:     model.Type,
		RefPrice: model.RefPrice,
		//Capacity:  model.Capacity,
	}
}

func HotelRoomsModelToDTO(models *[]model.HotelRoom) *[]HotelRoomDTO {
	var dtos []HotelRoomDTO
	for _, model := range *models {
		dtos = append(dtos, *HotelRoomModelToDTO(&model))
	}
	return &dtos
}
