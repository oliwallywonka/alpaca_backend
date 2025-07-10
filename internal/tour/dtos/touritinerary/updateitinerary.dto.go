package tidtos

import (
	"time"

	"github.com/oliwallywonka/alpaca_backend/db/model"
)

type UpdateItineraryDTO struct {
	Day int32 `json:"day"`
}

func (dto *UpdateItineraryDTO) DTOToModel() *model.Itinerary {
	return &model.Itinerary{
		Day:       dto.Day,
		UpdatedAt: time.Now().Unix(),
	}
}
