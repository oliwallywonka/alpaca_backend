package dtos

import "github.com/oliwallywonka/alpaca_backend/internal/departure/models"

type UpdateDepartureDTO struct {
	TourID         string `json:"tour_id"`
	StartDate      int64  `json:"start_date"`
	State          string `json:"state"`
	EndDate        int64  `json:"end_date"`
	AvailableSlots int    `json:"available_slots"`
}

func (dto *UpdateDepartureDTO) DTOToModel() *models.Departure {
	return &models.Departure{
		TourID:         dto.TourID,
		StartDate:      dto.StartDate,
		EndDate:        dto.EndDate,
		AvailableSlots: dto.AvailableSlots,
		State:          dto.State,
	}
}
