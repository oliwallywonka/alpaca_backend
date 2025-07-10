package tadtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/oliwallywonka/alpaca_backend/db/model"
)

type CreateTourActivityDTO struct {
	Day        int32  `json:"day" validate:"required"`
	Position   int32  `json:"position"`
	ActivityID string `json:"activityID" validate:"required"`
}

func (dto *CreateTourActivityDTO) DTOToModel(tourID string) *model.TourActivity {
	return &model.TourActivity{
		ID:         uuid.New().String(),
		TourID:     tourID,
		ActivityID: dto.ActivityID,
		Day:        dto.Day,
		Position:   dto.Position,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}
}
