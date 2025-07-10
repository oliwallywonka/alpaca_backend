package tmdtos

import (
	"time"

	"github.com/google/uuid"

	"github.com/oliwallywonka/alpaca_backend/db/model"
)

type CreateTourMealDTO struct {
	MealID string `json:"mealID" validate:"required"`
	Day    int32  `json:"day" validate:"gte=0"`
}

func (dto *CreateTourMealDTO) DTOToModel(tourID string) *model.TourMeals {
	return &model.TourMeals{
		ID:        uuid.New().String(),
		MealID:    dto.MealID,
		TourID:    tourID,
		Day:       dto.Day,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

}
