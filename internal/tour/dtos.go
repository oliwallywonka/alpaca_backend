package tourfx

import (
	"time"

	"github.com/google/uuid"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type CreateTourDTO struct {
	Code string
	Name shared.LangField
	Slug shared.LangField
}

func (dto *CreateTourDTO) DTOToModel() *model.Tour {
	return &model.Tour{
		ID:               uuid.NewString(),
		Code:             dto.Code,
		Name:             dto.Name,
		Slug:             dto.Slug,
		Transport:        shared.LangField{"en": ""},
		Accommodation:    shared.LangField{"en": ""},
		Team:             shared.LangField{"en": ""},
		ShortDescription: shared.LangField{"en": ""},
		LongDescription:  shared.LangField{"en": ""},
		Images:           shared.ImageField{},
		CreatedAt:        time.Now().Unix(),
		UpdatedAt:        time.Now().Unix(),
	}
}
