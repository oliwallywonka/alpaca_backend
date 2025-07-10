package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/oliwallywonka/alpaca_backend/internal/departure/models"
)

type CreateDepartureDTO struct {
	TourID         string `json:"tour_id" validate:"required"`
	StartDate      int64  `json:"start_date" validate:"required"`
	State          string `json:"state"`
	EndDate        int64  `json:"end_date"`
	AvailableSlots int    `json:"available_slots"`
}

func (dto *CreateDepartureDTO) DTOToModel() *models.Departure {
	return &models.Departure{
		ID:             uuid.New().String(),
		TourID:         dto.TourID,
		StartDate:      dto.StartDate,
		EndDate:        dto.EndDate,
		AvailableSlots: dto.AvailableSlots,
		CreatedAt:      time.Now().Unix(),
		UpdatedAt:      time.Now().Unix(),
		State:          dto.State,
	}
}