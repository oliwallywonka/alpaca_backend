package tddtos

import (
	"time"

	"github.com/oliwallywonka/alpaca_backend/db/model"
	destdtos "github.com/oliwallywonka/alpaca_backend/internal/destination/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/models"
)

type TourDestinationDTO struct {
	ID            string                  `json:"id"`
	DestinationID string                  `json:"destination_id"`
	TourID        string                  `json:"tour_id"`
	Description   shared.LangField        `json:"description"`
	Day           int32                   `json:"day"`
	Position      int32                   `json:"position"`
	Destination   destdtos.DestinationDTO `json:"destination,omitempty"`
}

func (dto *TourDestinationDTO) DTOToModel() *model.TourDestination {
	return &model.TourDestination{
		ID:            dto.ID,
		DestinationID: dto.DestinationID,
		TourID:        dto.TourID,
		Description:   dto.Description,
		Day:           dto.Day,
		Position:      dto.Position,
		UpdatedAt:     time.Now().Unix(),
	}
}

func TourDestinationModelToDTO(model *models.TourDestination) *TourDestinationDTO {

	dto := &TourDestinationDTO{
		ID:            model.ID,
		DestinationID: model.DestinationID,
		TourID:        model.TourID,
		Description:   model.Description,
		Day:           model.Day,
		Position:      model.Position,
		Destination:   *destdtos.DestinationModelToDTO(&model.Destination),
	}
	return dto
}

func TourDestinationsModelToDTO(models *[]models.TourDestination) *[]TourDestinationDTO {
	var dtos []TourDestinationDTO
	for _, model := range *models {
		dtos = append(dtos, *TourDestinationModelToDTO(&model))
	}
	return &dtos
}
