package repositories

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/hotel/models"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type Repository interface {
	GetPaginated(params *shared.PaginatedQueryParams) (*[]models.Hotel, int, error)
	GetByUniqueKey(key string) (*models.Hotel, error)
	SaveHotel(hotel *model.Hotel, rooms *[]model.HotelRoom) error
	UpdateHotel(hotel *model.Hotel, id string) error

	// HOTEL ROOMS
	GetHotelRooms(hotelID string) (*[]model.HotelRoom, error)
	CreateHotelRoom(hotelRooms *model.HotelRoom) error
	UpdateHotelRoom(hotelRoom *model.HotelRoom, hotelRoomID string) error
	DeleteHotelRoom(hotelRoomID string) error
}

type hotelRepository struct {
	db   *sql.DB
	gorm *gorm.DB
	ctx  context.Context
}

func New(db *sql.DB, gorm *gorm.DB, ctx context.Context) Repository {
	return &hotelRepository{
		db:   db,
		gorm: gorm,
		ctx:  ctx,
	}
}
