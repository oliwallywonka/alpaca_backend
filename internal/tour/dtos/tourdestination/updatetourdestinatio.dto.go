package tddtos

import (
	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type UpdateTourDestinationDTO struct {
	Description shared.LangField `json:"description"`
	Day         int32            `json:"day"`
	Position    int32            `json:"position"`
}

func (dto *UpdateTourDestinationDTO) DTOToModel() *model.TourDestination {
	return &model.TourDestination{
		Description: dto.Description,
		Day:         dto.Day,
		Position:    dto.Position,
	}
}
