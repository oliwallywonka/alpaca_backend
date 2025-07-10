package services

import (
	"github.com/oliwallywonka/alpaca_backend/internal/departure/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/departure/repositories"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type Service interface {
	GetDepartures(params *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error)
	GetDepartureByUniqueKey(key string) (*dtos.DepartureDTO, error)
	SaveDeparture(departure *dtos.CreateDepartureDTO) error
	UpdateDeparture(departureID string, departure *dtos.UpdateDepartureDTO) error
}

type departureService struct {
	rep repositories.Repository
}

func New(rep repositories.Repository) Service {
	return &departureService{
		rep: rep,
	}
}
