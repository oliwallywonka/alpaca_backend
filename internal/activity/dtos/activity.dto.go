package dtos

import (
	"github.com/oliwallywonka/alpaca_backend/internal/activity/models"
	destdtos "github.com/oliwallywonka/alpaca_backend/internal/destination/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type ActivityDTO struct {
	ID            string                  `json:"id"`
	DestinationID string                  `json:"destinationID"`
	Destination   destdtos.DestinationDTO `json:"destination,omitempty"`
	Title         shared.LangField        `json:"title"`
	Description   shared.LangField        `json:"description"`
	RefPrice      float64                 `json:"refPrice"`
	Images        []string                `json:"images"`
	CreatedAt     int64                   `json:"createdAt"`
	UpdatedAt     int64                   `json:"updatedAt"`
}

func ActivityModelToDTO(model *models.Activity) *ActivityDTO {
	dto := &ActivityDTO{
		ID:            model.ID,
		DestinationID: model.DestinationID,
		Title:         model.Title,
		Description:   model.Description,
		RefPrice:      model.RefPrice,
		Images:        model.Images,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
	}
	if model.Destination.ID != "" {
		dto.Destination = *destdtos.DestinationModelToDTO(&model.Destination)
	}
	return dto
}

func ActivityDTOsModelToDTO(models *[]models.Activity) *[]ActivityDTO {
	var dtos []ActivityDTO
	for _, model := range *models {
		dtos = append(dtos, *ActivityModelToDTO(&model))
	}
	return &dtos
}
