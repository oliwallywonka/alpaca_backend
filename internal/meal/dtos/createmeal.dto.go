package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type CreateMealDTO struct {
	Type        string           `json:"type" validate:"required,oneof=breakfast lunch dinner extra"`
	RefPrice    float64          `json:"refPrice"`
	Description shared.LangField `json:"description"`
}

func (dto *CreateMealDTO) DTOToModel() *model.Meal {
	return &model.Meal{
		ID:          uuid.New().String(),
		Type:        dto.Type,
		RefPrice:    dto.RefPrice,
		Description: dto.Description,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
}
