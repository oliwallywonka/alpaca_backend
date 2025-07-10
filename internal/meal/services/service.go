package services

import (
	"github.com/oliwallywonka/alpaca_backend/internal/meal/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/meal/repositories"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type Service interface {
	GetPaginated(dto *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error)
	GetByUniqueKey(key string) (*dtos.MealDTO, error)
	Save(meal *dtos.CreateMealDTO) error
	Update(mealID string, meal *dtos.UpdateMealDTO) error
}

type mealService struct {
	rep repositories.Repository
}

func New(rep repositories.Repository) Service {
	return &mealService{
		rep: rep,
	}
}
