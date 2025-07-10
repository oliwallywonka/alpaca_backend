package dtos

import (
	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type UpdateActivityDTO struct {
	DestinationID string           `json:"destinationID"`
	Title         shared.LangField `json:"title"`
	Description   shared.LangField `json:"description"`
	RefPrice      float64          `json:"refPrice"`
}

func (dto *UpdateActivityDTO) DTOToModel() *model.Activity {
	return &model.Activity{
		Title:         dto.Title,
		DestinationID: dto.DestinationID,
		Description:   dto.Description,
		RefPrice:      dto.RefPrice,
	}
}
