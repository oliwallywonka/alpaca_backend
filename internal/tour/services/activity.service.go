package services

import (
	tadtos "github.com/oliwallywonka/alpaca_backend/internal/tour/dtos/touractivity"
)

func (s *tourService) CreateTourActivity(tourID string, dto *tadtos.CreateTourActivityDTO) error {
	_, err := s.rep.GetByUniqueKey(tourID)
	if err != nil {
		return err
	}

	_, err = s.actRep.GetByUniqueKey(dto.ActivityID)
	if err != nil {
		return err
	}

	err = s.rep.CreateTourActivity(dto.DTOToModel(tourID))
	if err != nil {
		return err
	}
	return nil
}

func (s *tourService) UpdateTourActivity(tourActivityID string, dto *tadtos.UpdateTourActivityDTO) error {
	_, err := s.rep.GetTourActivityByID(tourActivityID)
	if err != nil {
		return err
	}

	err = s.rep.UpdateTourActivity(tourActivityID, dto.DTOToModel())
	if err != nil {
		return err
	}
	return nil
}

func (s *tourService) GetTourActivities(tourID string) (*[]tadtos.TourActivityDTO, error) {
	_, err := s.rep.GetByUniqueKey(tourID)
	if err != nil {
		return nil, err
	}
	
	activities, err := s.rep.GetTourActivities(tourID)
	if err != nil {
		return nil, err
	}
	return tadtos.TourActivityModelsToDTO(activities), nil
}

func (s *tourService) DeleteTourActivity(tourActivityID string) error {
	err := s.rep.DeleteTourActivity(tourActivityID)
	if err != nil {
		return err
	}
	return nil
}