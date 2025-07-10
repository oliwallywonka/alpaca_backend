package services

import (
	"fmt"

	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/tourerrors"
)

func (s *tourService) NameOrSlugExists(name string) (bool, error) {
	return s.rep.NameOrSlugExists(name)
}

func (s *tourService) SaveTour(dto *dtos.CreateTourDTO) (tourID *string, err error) {

	for _, name := range dto.Name {
		tourName, _ := s.rep.NameOrSlugExists(name)
		if tourName {
			return nil, tourerrors.UniqueNameError
		}
	}

	for _, slug := range dto.Slug {
		tourSlug, _ := s.rep.NameOrSlugExists(slug)
		if tourSlug {
			return nil, tourerrors.UniqueSlugError
		}
	}

	tourID, err = s.rep.SaveTour(dto.DTOToModel())
	if err != nil {
		return nil, err
	}
	return tourID, nil
}

func (s *tourService) GetByUniqueKey(slug string) (*dtos.TourDTO, error) {
	tour, err := s.rep.GetByUniqueKey(slug)
	if err != nil {
		return nil, err
	}
	return dtos.TourModelToDTO(tour), nil
}

func (s *tourService) GetTotalCount() int {
	return 0
}

func (s *tourService) GetPaginated(dto *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error) {
	tours, totalTours, err := s.rep.GetPaginated(dto.DTOToModel())
	if err != nil {
		return nil, err
	}
	total := totalTours
	page := dto.Page
	response := shared.PaginatedResponse{Items: dtos.ToursModelToDTO(tours), Total: total, Page: int(page)}
	return &response, nil
}

func (s *tourService) UpdateTour(dto *dtos.UpdateTourDTO, id string) error {

	fmt.Printf("%+v\n", dto)
	_, err := s.rep.GetByUniqueKey(id)
	if err != nil {
		return err
	}

	// TODO: compare if name and slug are in the found tour before returning a unique error
	/*if dto.Name != nil {
		for _, name := range dto.Name {
			tourName, _ := s.rep.NameExists(name)
			if tourName {
				return tourerrors.UniqueNameError
			}
		}
	}

	if dto.Slug != nil {
		for _, slug := range dto.Slug {
			tourSlug, _ := s.rep.SlugExists(slug)
			if tourSlug {
				return tourerrors.UniqueSlugError
			}
		}
	}*/

	err = s.rep.UpdateTour(id, dto.DTOToModel())
	if err != nil {
		return err
	}
	return nil
}
