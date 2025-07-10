package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type CreateActivityDTO struct {
	Title         shared.LangField `json:"title" validate:"required"`
	DestinationID string           `json:"destinationID" validate:"required"`
	Description   shared.LangField `json:"description"`
	RefPrice      float64          `json:"refPrice"`
}

func (dto *CreateActivityDTO) ToModel() *model.Activity {
	return &model.Activity{
		ID:            uuid.New().String(),
		Title:         dto.Title,
		DestinationID: dto.DestinationID,
		Description:   dto.Description,
		RefPrice:      dto.RefPrice,
		CreatedAt:     time.Now().Unix(),
		UpdatedAt:     time.Now().Unix(),
	}
}
