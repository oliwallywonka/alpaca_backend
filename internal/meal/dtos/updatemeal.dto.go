package dtos

import (
	"time"

	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type UpdateMealDTO struct {
	Type        string           `json:"type" validate:"oneof=breakfast lunch dinner extra"`
	RefPrice    float64          `json:"refPrice"`
	Description shared.LangField `json:"description"`
}

func (dto *UpdateMealDTO) DTOToModel() *model.Meal {
	return &model.Meal{
		Type:        dto.Type,
		RefPrice:    dto.RefPrice,
		Description: dto.Description,
		UpdatedAt:   time.Now().Unix(),
	}
}
