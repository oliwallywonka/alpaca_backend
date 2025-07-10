package dtos

import (
	"github.com/oliwallywonka/alpaca_backend/db/model"
)

type UpdateHotelDTO struct {
	Name      string `json:"name"`
	Direction string `json:"direction"`
	Phone     string `json:"phone"`
}

func (dto *UpdateHotelDTO) DTOToModel() *model.Hotel {
	return &model.Hotel{
		Name:      dto.Name,
		Direction: dto.Direction,
		Phone:     dto.Phone,
	}
}