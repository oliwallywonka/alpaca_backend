package services

import (
	"github.com/oliwallywonka/alpaca_backend/internal/hotel/repositories"
	"github.com/oliwallywonka/alpaca_backend/internal/hotel/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type Service interface {
	GetPaginated(dto *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error)
	GetByUniqueKey(key string) (*dtos.HotelDTO, error)
	SaveHotel(dto *dtos.CreateHotelDTO) error
	UpdateHotel(dto *dtos.UpdateHotelDTO, id string) error

	// HOTEL ROOMS
	CreateHotelRoom(dto *dtos.CreateHotelRoomDTO, hotelID string) error
	UpdateHotelRoom(dto *dtos.UpdateHotelRoomDTO, hotelRoomID string) error
	GetHotelRooms(hotelID string) (*[]dtos.HotelRoomDTO, error)
	DeleteHotelRoom(hotelRoomID string) error
}

type hotelService struct {
	rep repositories.Repository
}

func New(rep repositories.Repository) Service {
	return &hotelService{
		rep: rep,
	}
}

