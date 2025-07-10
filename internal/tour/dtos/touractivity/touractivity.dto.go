package tadtos

import (
	adtos "github.com/oliwallywonka/alpaca_backend/internal/activity/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/models"
)

type TourActivityDTO struct {
	ID         string             `json:"id"`
	TourID     string             `json:"tour_id"`
	ActivityID string             `json:"activity_id"`
	Day        int32              `json:"day"`
	Position   int32              `json:"position"`
	CreatedAt  int64              `json:"created_at"`
	UpdatedAt  int64              `json:"updated_at"`
	Activity   *adtos.ActivityDTO `json:"activity,omitempty"`
}

func TourActivityModelToDTO(model *models.TourActivity) *TourActivityDTO {
	dto := &TourActivityDTO{
		ID:         model.ID,
		TourID:     model.TourID,
		ActivityID: model.ActivityID,
		Day:        model.Day,
		Position:   model.Position,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
		Activity:   adtos.ActivityModelToDTO(&model.Activity),
	}
	return dto
}

func TourActivityModelsToDTO(models *[]models.TourActivity) *[]TourActivityDTO {
	var dtos []TourActivityDTO
	for _, model := range *models {
		dtos = append(dtos, *TourActivityModelToDTO(&model))
	}
	return &dtos
}
