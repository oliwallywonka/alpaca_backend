package services

import (
	tmdtos "github.com/oliwallywonka/alpaca_backend/internal/tour/dtos/tourmeal"
)

func (s *tourService) CreateTourMeal(tourID string, dto *tmdtos.CreateTourMealDTO) error {
	_, err := s.mealRep.GetByUniqueKey(dto.MealID)
	if err != nil {
		return err
	}
	_, err = s.rep.GetByUniqueKey(tourID)
	if err != nil {
		return err
	}
	err = s.rep.CreateTourMeal(dto.DTOToModel(tourID))
	if err != nil {
		return err
	}
	return nil
}

func (s *tourService) UpdateTourMeal(tourMealID string, dto *tmdtos.UpdateTourMealDTO) error {
	_, err := s.rep.GetTourMealByID(tourMealID)
	if err != nil {
		return err
	}

	err = s.rep.UpdateTourMeal(tourMealID, dto.DTOToModel())
	if err != nil {
		return err
	}
	return nil
}

func (s *tourService) GetTourMeals(tourID string) (*[]tmdtos.TourMealDTO, error) {
	_, err := s.rep.GetByUniqueKey(tourID)
	if err != nil {
		return nil, err
	}
	meals, err := s.rep.GetTourMeals(tourID)
	if err != nil {
		return nil, err
	}
	return tmdtos.TourMealModelsToDTO(meals), nil
}

func (s *tourService) DeleteTourMeal(tourMealID string) error {
	err := s.rep.DeleteTourMeal(tourMealID)
	if err != nil {
		return err
	}
	return nil
}
