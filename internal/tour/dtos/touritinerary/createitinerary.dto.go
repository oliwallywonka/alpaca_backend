package tidtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/oliwallywonka/alpaca_backend/db/model"
)

type CreateItineraryDTO struct {
	Day int32 `json:"day"`
}

func (dto *CreateItineraryDTO) DTOToModel(tourID string) *model.Itinerary {
	return &model.Itinerary{
		ID:        uuid.New().String(),
		Day:       dto.Day,
		TourID:    tourID,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}
