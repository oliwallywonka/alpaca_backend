package repositories

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/activity/models"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type Repository interface {
	GetPaginated(params *shared.PaginatedQueryParams) (*[]models.Activity, int, error)
	GetByUniqueKey(key string) (*models.Activity, error)
	Create(activity *model.Activity) error
	Update(activity *model.Activity, id string) error
}

type activityRepository struct {
	db   *sql.DB
	gorm *gorm.DB
	ctx  context.Context
}

func New(db *sql.DB, gorm *gorm.DB, ctx context.Context) Repository {
	return &activityRepository{
		db:   db,
		gorm: gorm,
		ctx:  ctx,
	}
}
