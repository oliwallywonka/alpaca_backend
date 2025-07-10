package tmdtos

import (
	"github.com/oliwallywonka/alpaca_backend/db/model"
)

type UpdateTourMealDTO struct {
	Day int32 `json:"day"`
}

func (dto *UpdateTourMealDTO) DTOToModel() *model.TourMeals {
	return &model.TourMeals{
		Day: dto.Day,
	}
}
