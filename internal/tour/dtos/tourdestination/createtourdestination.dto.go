package tddtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type CreateTourDestinationDTO struct {
	DestinationID string           `json:"destinationID"`
	Description   shared.LangField `json:"description" validate:"required"`
	Day           int32            `json:"day" validate:"gte=0"`
	Position      int32            `json:"position"`
}

func (dto *CreateTourDestinationDTO) DTOToModel(tourID string) *model.TourDestination {
	return &model.TourDestination{
		ID:            uuid.New().String(),
		TourID:        tourID,
		DestinationID: dto.DestinationID,
		Description:   dto.Description,
		Day:           dto.Day,
		Position:      dto.Position,
		CreatedAt:     time.Now().Unix(),
		UpdatedAt:     time.Now().Unix(),
	}
}
