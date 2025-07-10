package dtos

import "github.com/oliwallywonka/alpaca_backend/internal/hotel/models"

type HotelDTO struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Direction string         `json:"direction"`
	Phone     string         `json:"phone"`
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
	Rooms     []HotelRoomDTO `json:"rooms"`
}

func HotelModelToDTO(model *models.Hotel) *HotelDTO {
	return &HotelDTO{
		ID:        model.ID,
		Name:      model.Name,
		Direction: model.Direction,
		Phone:     model.Phone,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Rooms:     *HotelRoomsModelToDTO(&model.Rooms),
	}
}

func HotelDTOsModelToDTO(models *[]models.Hotel) *[]HotelDTO {
	var dtos []HotelDTO
	for _, model := range *models {
		dtos = append(dtos, *HotelModelToDTO(&model))
	}
	return &dtos
}
