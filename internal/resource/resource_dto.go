package resourcefx

import (
	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)



type UpdateResourceDTO struct {
	Name           shared.LangField
	Description    shared.LangField
	resourceTypeID string
	Location       shared.LocationField
	Images         shared.ImageField
	IsActive       *bool
	UpdatedAt      int
}

type ResourceProviderDTO struct {
	model.ResourceProvider
	Provider model.Provider
	User     model.User
	Resource model.Resource
}
