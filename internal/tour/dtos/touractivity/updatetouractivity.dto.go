package tadtos

import (
	"time"

	"github.com/oliwallywonka/alpaca_backend/db/model"
)

type UpdateTourActivityDTO struct {
	Day      int32 `json:"day"`
	Position int32 `json:"position"`
}

func (dto *UpdateTourActivityDTO) DTOToModel() *model.TourActivity {
	return &model.TourActivity{
		Day:       dto.Day,
		Position:  dto.Position,
		UpdatedAt: time.Now().Unix(),
	}
}
