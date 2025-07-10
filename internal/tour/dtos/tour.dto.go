package dtos

import (
	destdtos "github.com/oliwallywonka/alpaca_backend/internal/destination/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/models"
)

type TourDTO struct {
	ID               string                    `json:"id"`
	Name             shared.LangField          `json:"name"`
	Slug             shared.LangField          `json:"slug"`
	RefPricePP       float64                   `json:"refPricePP"`
	Days             int32                     `json:"days"`
	GroupSize        string                    `json:"groupSize"`
	Transport        shared.LangField          `json:"transport"`
	Accommodation    shared.LangField          `json:"accommodation"`
	Team             shared.LangField          `json:"team"`
	ShortDescription shared.LangField          `json:"shortDescription"`
	LongDescription  shared.LangField          `json:"longDescription"`
	Images           shared.ImageField         `json:"images"`
	Destinations     []destdtos.DestinationDTO `json:"destinations"`
	IsPublic         bool                      `json:"isPublic"`
	CreatedAt        int64                     `json:"createdAt"`
	UpdatedAt        int64                     `json:"updatedAt"`
}

func TourModelToDTO(model *models.Tour) *TourDTO {
	return &TourDTO{
		ID:               model.ID,
		Name:             model.Name,
		Slug:             model.Slug,
		RefPricePP:       model.RefPricePp,
		Days:             model.Days,
		GroupSize:        model.GroupSize,
		Transport:        model.Transport,
		Accommodation:    model.Accommodation,
		Team:             model.Team,
		ShortDescription: model.ShortDescription,
		LongDescription:  model.LongDescription,
		Images:           model.Images,
		IsPublic:         model.IsPublic,
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
		Destinations:     *destdtos.DestinationModelsToDTO(&model.Destinations),
	}
}

func ToursModelToDTO(models *[]models.Tour) *[]TourDTO {
	var dtos []TourDTO
	for _, model := range *models {
		dtos = append(dtos, *TourModelToDTO(&model))
	}
	return &dtos
}
