package repositories

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type Repository interface {
	GetByUniqueKey(key string) (*model.Meal, error)
	GetPaginated(params *shared.PaginatedQueryParams) (*[]model.Meal, int, error)
	Save(meal *model.Meal) error
	Update(mealID string, meal *model.Meal) error
}

type mealRepository struct {
	db *sql.DB
	gorm *gorm.DB
}

func New(db *sql.DB, gorm *gorm.DB) Repository {
	return &mealRepository{
		db: db,
		gorm: gorm,
	}
}
