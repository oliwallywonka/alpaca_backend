package repositories

import (
	"context"
	"database/sql"

	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"gorm.io/gorm"
)

type Repository interface {
	GetPaginated(*shared.PaginatedQueryParams) (*[]model.Destination, int, error)
	GetByUniqueKey(string) (*[]model.Destination, error)
	NameExists(string) (bool, error)
	Create(destination *model.Destination) (*model.Destination, error)
	Update(id string, destination *model.Destination) (*model.Destination, error)
}

type destinationRepository struct {
	db   *sql.DB
	gorm *gorm.DB
	ctx  context.Context
}

func New(db *sql.DB, gorm *gorm.DB, ctx context.Context) Repository {
	return &destinationRepository{
		db:   db,
		gorm: gorm,
		ctx:  ctx,
	}
}
