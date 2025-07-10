package services

import (
	"github.com/oliwallywonka/alpaca_backend/internal/meal/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

func (s *mealService) GetPaginated(dto *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error) {
	meals, total, err := s.rep.GetPaginated(dto.DTOToModel())
	if err != nil {
		return nil, err
	}
	dtos := dtos.MealDTOsModelToDTO(meals)
	return &shared.PaginatedResponse{Items: dtos, Total: total, Page: int(dto.Page)}, nil
}

func (s *mealService) GetByUniqueKey(key string) (*dtos.MealDTO, error) {
	meal, err := s.rep.GetByUniqueKey(key)
	if err != nil {
		return nil, err
	}
	return dtos.MealModelToDTO(meal), nil
}

func (s *mealService) Save(meal *dtos.CreateMealDTO) error {
	err := s.rep.Save(meal.DTOToModel())
	if err != nil {
		return err
	}
	return nil
}

func (s *mealService) Update(mealID string, meal *dtos.UpdateMealDTO) error {
	_, err := s.rep.GetByUniqueKey(mealID)
	if err != nil {
		return err
	}
	err = s.rep.Update(mealID, meal.DTOToModel())
	if err != nil {
		return err
	}
	return nil
}
