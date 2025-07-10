package services

import "github.com/oliwallywonka/alpaca_backend/internal/hotel/dtos"

func (s *hotelService) GetHotelRooms(hotelID string) (*[]dtos.HotelRoomDTO, error) {
	hotelRooms, err := s.rep.GetHotelRooms(hotelID)
	if err != nil {
		return nil, err
	}
	return dtos.HotelRoomsModelToDTO(hotelRooms), nil
}

func (s *hotelService) CreateHotelRoom(dto *dtos.CreateHotelRoomDTO, hotelID string) error {
	hotelRoom := dto.DTOToModel(hotelID)
	err := s.rep.CreateHotelRoom(hotelRoom)
	if err != nil {
		return err
	}
	return nil
}

func (s *hotelService) UpdateHotelRoom(dto *dtos.UpdateHotelRoomDTO, hotelRoomID string) error {
	_, err := s.rep.GetByUniqueKey(hotelRoomID)
	if err != nil {
		return err
	}
	err = s.rep.UpdateHotelRoom(dto.DTOToModel(), hotelRoomID)
	if err != nil {
		return err
	}
	return nil
}

func (s *hotelService) DeleteHotelRoom(hotelRoomID string) error {
	return s.rep.DeleteHotelRoom(hotelRoomID)
}
