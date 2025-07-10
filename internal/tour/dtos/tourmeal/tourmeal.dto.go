package tmdtos

import (
	mdtos "github.com/oliwallywonka/alpaca_backend/internal/meal/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/models"
)

type TourMealDTO struct {
	ID        string         `json:"id"`
	TourID    string         `json:"tour_id"`
	MealID    string         `json:"meal_id"`
	Day       int32          `json:"day"`
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
	Meal      *mdtos.MealDTO `json:"meal,omitempty"`
}

func TourMealModelToDTO(model *models.TourMeal) *TourMealDTO {
	dto := &TourMealDTO{
		ID:        model.ID,
		TourID:    model.TourID,
		MealID:    model.MealID,
		Day:       model.Day,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Meal:      mdtos.MealModelToDTO(&model.Meal),
	}
	return dto
}

func TourMealModelsToDTO(models *[]models.TourMeal) *[]TourMealDTO {
	var dtos []TourMealDTO
	for _, model := range *models {
		dtos = append(dtos, *TourMealModelToDTO(&model))
	}
	return &dtos
}
