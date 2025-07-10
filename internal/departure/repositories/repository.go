package repositories

import (
	"github.com/oliwallywonka/alpaca_backend/internal/departure/models"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"gorm.io/gorm"
)

type Repository interface {
	GetDepartures(params *shared.PaginatedQueryParams) ([]*models.Departure, int, error)
	GetDepartureByUniqueKey(key string) (*models.Departure, error)
	SaveDeparture(departure *models.Departure) error
	UpdateDeparture(departureID string, departure *models.Departure) error
}

type departureRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &departureRepository{
		db: db,
	}
}
