package dtos

import (
	"github.com/oliwallywonka/alpaca_backend/db/model"
)

type UpdateHotelRoomDTO struct {
	RefPrice float64 `json:"ref_price"`
	Type     string  `json:"type" validate:"oneof=single double triple quadruple"`
}

func (dto *UpdateHotelRoomDTO) DTOToModel() *model.HotelRoom {
	return &model.HotelRoom{
		RefPrice: dto.RefPrice,
		Type:     dto.Type,
	}
}
