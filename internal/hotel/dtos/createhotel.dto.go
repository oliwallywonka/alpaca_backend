package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/oliwallywonka/alpaca_backend/db/model"
)

type CreateHotelDTO struct {
	Name      string                `json:"name" validate:"required"`
	Direction string                `json:"direction" validate:"required"`
	Phone     string                `json:"phone" validate:"required"`
	Rooms     *[]CreateHotelRoomDTO `json:"rooms"`
}

func (dto *CreateHotelDTO) DTOToModel() *model.Hotel {
	return &model.Hotel{
		ID:        uuid.New().String(),
		Name:      dto.Name,
		Direction: dto.Direction,
		Phone:     dto.Phone,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}
