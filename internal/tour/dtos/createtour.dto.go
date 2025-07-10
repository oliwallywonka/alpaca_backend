package dtos

import (
	"time"

	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type CreateTourDTO struct {
	Name shared.LangField `json:"name" validate:"required"`
	Slug shared.LangField `json:"slug" validate:"required"`
}

func (dto *CreateTourDTO) DTOToModel() *model.Tour {
	return &model.Tour{
		Name:     dto.Name,
		Slug:     dto.Slug,
		Images:   nil,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}
