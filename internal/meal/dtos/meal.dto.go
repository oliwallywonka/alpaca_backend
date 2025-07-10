package dtos

import (
	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type MealDTO struct {
	ID          string           `json:"id"`
	Type        string           `json:"type"`
	RefPrice    float64          `json:"refPrice"`
	Description shared.LangField `json:"description"`
	CreatedAt   int64            `json:"createdAt"`
	UpdatedAt   int64            `json:"updatedAt"`
}

func MealModelToDTO(model *model.Meal) *MealDTO {
	return &MealDTO{
		ID:          model.ID,
		Type:        model.Type,
		RefPrice:    model.RefPrice,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func MealDTOsModelToDTO(models *[]model.Meal) *[]MealDTO {
	var dtos []MealDTO
	for _, model := range *models {
		dtos = append(dtos, *MealModelToDTO(&model))
	}
	return &dtos
}
