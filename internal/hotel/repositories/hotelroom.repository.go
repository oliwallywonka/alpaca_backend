package repositories

import (
	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	. "github.com/oliwallywonka/alpaca_backend/db/table"
)

func (r *hotelRepository) GetHotelRooms(hotelID string) (*[]model.HotelRoom, error) {
	var hotelRooms []model.HotelRoom
	smt := HotelRoom.SELECT(HotelRoom.AllColumns).
		FROM(HotelRoom).
		WHERE(HotelRoom.HotelID.EQ(String(hotelID)))
	err := smt.Query(r.db, &hotelRooms)
	if err != nil {
		return nil, err
	}
	return &hotelRooms, nil
}

func (r *hotelRepository) CreateHotelRoom(hotelRoom *model.HotelRoom) error {
	stm := HotelRoom.INSERT(HotelRoom.AllColumns).
		MODEL(hotelRoom)
	_, err := stm.Exec(r.db)
	if err != nil {
		return err
	}
	return nil
}

func (r *hotelRepository) UpdateHotelRoom(hotelRoom *model.HotelRoom, hotelRoomID string) error {
	return r.gorm.
		Table("hotel_room").
		Where("id = ?", hotelRoomID).
		Updates(hotelRoom).
		Error
}

func (r *hotelRepository) DeleteHotelRoom(hotelRoomID string) error {
	stm := HotelRoom.DELETE().
		WHERE(HotelRoom.ID.EQ(String(hotelRoomID)))
	_, err := stm.Exec(r.db)
	if err != nil {
		return err
	}
	return nil
}
