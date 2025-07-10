package services

import (
	"fmt"

	tddtos "github.com/oliwallywonka/alpaca_backend/internal/tour/dtos/tourdestination"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/tourerrors"
)

func (s *tourService) CreateTourDestination(dto *tddtos.CreateTourDestinationDTO, tourID string) error {
	fmt.Printf("%+v\n", dto)
	fmt.Printf("%+v\n", tourID)
	_, err := s.rep.GetByUniqueKey(tourID)
	if err != nil {
		return err
	}
	_, err = s.destRep.GetByUniqueKey(dto.DestinationID)
	if err != nil {
		return err
	}

	err = s.rep.CreateTourDestination(dto.DTOToModel(tourID))
	if err != nil {
		return err
	}
	return nil
}

func (s *tourService) UpdateTourDestination(dto *tddtos.UpdateTourDestinationDTO, tourID string, tourDestinationID string) error {
	_, err := s.rep.GetByUniqueKey(tourID)
	if err != nil {
		return err
	}

	_, err = s.rep.GetTourDestinationByID(tourDestinationID)

	if err != nil {
		return err
	}

	err = s.rep.UpdateTourDestination(dto.DTOToModel(), tourDestinationID)
	if err != nil {
		return err
	}
	return nil
}

func (s *tourService) GetTourDestinations(tourID string) (*[]tddtos.TourDestinationDTO, error) {
	_, err := s.rep.GetByUniqueKey(tourID)
	if err != nil {
		return nil, tourerrors.NotFoundError
	}
	tourDestinations, err := s.rep.GetTourDestinations(tourID)
	if err != nil {
		return nil, err
	}
	dtos := tddtos.TourDestinationsModelToDTO(tourDestinations)
	return dtos, nil
}

func (s *tourService) DeleteTourDestination(destinationID string) error {
	err := s.rep.DeleteTourDestination(destinationID)
	if err != nil {
		return err
	}
	return nil
}
