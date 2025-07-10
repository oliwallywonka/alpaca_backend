package dtos

import (
	"github.com/oliwallywonka/alpaca_backend/internal/departure/models"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type DepartureDTO struct {
	ID             string            `json:"id"`
	TourID         string            `json:"tour_id"`
	State          string            `json:"state"`
	StartDate      int64             `json:"start_date"`
	EndDate        int64             `json:"end_date"`
	AvailableSlots int               `json:"available_slots"`
	CreatedAt      int64             `json:"created_at"`
	UpdatedAt      int64             `json:"updated_at"`
	Tour           *TourDepartureDTO `json:"tour"`
}

type TourDepartureDTO struct {
	Name       shared.LanguageField `json:"name"`
	RefPricePP float32              `json:"ref_price_pp"`
	Days       int8                 `json:"days"`
	Images     shared.Images        `json:"images"`
}

func DepartureModelToDTO(departure *models.Departure) *DepartureDTO {
	return &DepartureDTO{
		ID:             departure.ID,
		TourID:         departure.TourID,
		State:          departure.State,
		StartDate:      departure.StartDate,
		EndDate:        departure.EndDate,
		AvailableSlots: departure.AvailableSlots,
		CreatedAt:      departure.CreatedAt,
		UpdatedAt:      departure.UpdatedAt,
		Tour:           TourDepartureModelToDTO(departure.Tour),
	}
}
func TourDepartureModelToDTO(tourDeparture *models.TourDeparture) *TourDepartureDTO {
	return &TourDepartureDTO{
		Name:       tourDeparture.Name,
		RefPricePP: tourDeparture.RefPricePP,
		Days:       tourDeparture.Days,
		Images:     tourDeparture.Images,
	}
}

func DepartureModelsToDTO(departures []*models.Departure) []*DepartureDTO {
	var dtos []*DepartureDTO
	for _, departure := range departures {
		dtos = append(dtos, DepartureModelToDTO(departure))
	}
	return dtos
}
