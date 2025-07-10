package services

import (
	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/hotel/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

func (s *hotelService) GetPaginated(dto *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error) {
	hotels, totalHotels, err := s.rep.GetPaginated(dto.DTOToModel())
	if err != nil {
		return nil, err
	}
	response := shared.PaginatedResponse{Items: dtos.HotelDTOsModelToDTO(hotels), Total: totalHotels, Page: int(dto.Page)}
	return &response, nil
}

func (s *hotelService) GetByUniqueKey(key string) (*dtos.HotelDTO, error) {
	hotel, err := s.rep.GetByUniqueKey(key)
	if err != nil {
		return nil, err
	}
	return dtos.HotelModelToDTO(hotel), nil
}

func (s *hotelService) SaveHotel(dto *dtos.CreateHotelDTO) error {
	hotel := dto.DTOToModel()
	var rooms []model.HotelRoom
	if dto.Rooms != nil {
		for _, dto := range *dto.Rooms {
			room := dto.DTOToModel(hotel.ID)
			rooms = append(rooms, *room)
		}
	}
	err := s.rep.SaveHotel(hotel, &rooms)
	if err != nil {
		return err
	}
	return nil
}

func (s *hotelService) UpdateHotel(dto *dtos.UpdateHotelDTO, id string) error {
	err := s.rep.UpdateHotel(dto.DTOToModel(), id)
	if err != nil {
		return err
	}
	return nil
}
