package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/oliwallywonka/alpaca_backend/db/model"
)

type CreateHotelRoomDTO struct {
	Type     string  `json:"type" validate:"required,oneof=single double shared"`
	RefPrice float64 `json:"ref_price"`
	//Capacity int     `json:"capacity"`
}

func (dto *CreateHotelRoomDTO) DTOToModel(hotelID string) *model.HotelRoom {
	return &model.HotelRoom{
		ID:        uuid.New().String(),
		HotelID:   hotelID,
		Type:      dto.Type,
		RefPrice:  dto.RefPrice,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}
