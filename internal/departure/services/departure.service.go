package services

import (
	"github.com/oliwallywonka/alpaca_backend/internal/departure/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

func (s *departureService) GetDepartureByUniqueKey(key string) (*dtos.DepartureDTO, error) {
	departure, err := s.rep.GetDepartureByUniqueKey(key)
	if err != nil {
		return nil, err
	}
	return dtos.DepartureModelToDTO(departure), nil
}

func (s *departureService) GetDepartures(params *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error) {
	departures, total, err := s.rep.GetDepartures(params.DTOToModel())
	if err != nil {
		return nil, err
	}
	page := params.Page
	response := shared.PaginatedResponse{Items: dtos.DepartureModelsToDTO(departures), Total: total, Page: page}
	return &response, nil
}

func (s *departureService) SaveDeparture(departure *dtos.CreateDepartureDTO) error {
	err := s.rep.SaveDeparture(departure.DTOToModel())
	if err != nil {
		return err
	}
	return nil
}

func (s *departureService) UpdateDeparture(departureID string, departure *dtos.UpdateDepartureDTO) error {
	_, err := s.rep.GetDepartureByUniqueKey(departureID)
	if err != nil {
		return err
	}
	err = s.rep.UpdateDeparture(departureID, departure.DTOToModel())
	if err != nil {
		return err
	}
	return nil
}