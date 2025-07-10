package dtos

import (
	"time"

	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type UpdateTourDTO struct {
	Name             shared.LangField `json:"name"`
	Slug             shared.LangField `json:"slug"`
	RefPricePP       float64          `json:"refPricePP"`
	Days             int32            `json:"days"`
	GroupSize        string           `json:"groupSize"`
	Transport        shared.LangField `json:"transport"`
	Accommodation    shared.LangField `json:"accommodation"`
	Team             shared.LangField `json:"team"`
	ShortDescription shared.LangField `json:"shortDescription"`
	LongDescription  shared.LangField `json:"longDescription"`
	IsPublic         bool             `json:"isPublic"`
}

func (dto *UpdateTourDTO) DTOToModel() *model.Tour {
	return &model.Tour{
		Name:             dto.Name,
		Slug:             dto.Slug,
		RefPricePp:       dto.RefPricePP,
		Days:             dto.Days,
		GroupSize:        dto.GroupSize,
		Transport:        dto.Transport,
		Accommodation:    dto.Accommodation,
		Team:             dto.Team,
		ShortDescription: dto.ShortDescription,
		LongDescription:  dto.LongDescription,
		IsPublic:         dto.IsPublic,
		UpdatedAt:        time.Now().Unix(),
	}
}
