package tidtos

import (
	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type ItineraryDTO struct {
	ID        string            `json:"id"`
	TourID    string            `json:"tour_id"`
	Images    shared.ImageField `json:"images"`
	Day       int32             `json:"day"`
	CreatedAt int64             `json:"created_at"`
	UpdatedAt int64             `json:"updated_at"`
}

func ItineraryModelToDTO(model *model.Itinerary) *ItineraryDTO {
	return &ItineraryDTO{
		ID:        model.ID,
		TourID:    model.TourID,
		Images:    model.Images,
		Day:       model.Day,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func ItineraryDTOsModelToDTO(models *[]model.Itinerary) *[]ItineraryDTO {
	var dtos []ItineraryDTO
	for _, model := range *models {
		dtos = append(dtos, *ItineraryModelToDTO(&model))
	}
	return &dtos
}
