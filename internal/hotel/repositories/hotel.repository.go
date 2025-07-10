package repositories

import (
	"errors"
	"fmt"

	"github.com/go-jet/jet/v2/qrm"
	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	. "github.com/oliwallywonka/alpaca_backend/db/table"
	"github.com/oliwallywonka/alpaca_backend/internal/hotel/hotelerrors"
	"github.com/oliwallywonka/alpaca_backend/internal/hotel/models"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

func (r *hotelRepository) GetPaginated(params *shared.PaginatedQueryParams) (*[]models.Hotel, int, error) {
	var countStruct []struct {
		Total int
	}
	var hotels []models.Hotel

	searchFilter := fmt.Sprintf("%%%s%%", params.SearchFilter)
	smt := Hotel.SELECT(Hotel.AllColumns, HotelRoom.AllColumns).
		FROM(Hotel.LEFT_JOIN(HotelRoom, Hotel.ID.EQ(HotelRoom.HotelID))).
		WHERE(OR(
			CAST(Hotel.Name).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Hotel.Direction).AS_TEXT().LIKE(String(searchFilter)),
		)).
		LIMIT(params.Limit).
		OFFSET(params.Offset)
	err := smt.Query(r.db, &hotels)
	if err != nil {
		return nil, 0, err
	}

	smtCount := SELECT(COUNT(STAR).AS("total")).
		FROM(Hotel).
		WHERE(OR(
			CAST(Hotel.Name).AS_TEXT().LIKE(String(params.SearchFilter)),
			CAST(Hotel.Direction).AS_TEXT().LIKE(String(params.SearchFilter)),
		))
	err = smtCount.Query(r.db, &countStruct)
	if err != nil {
		return nil, 0, err
	}
	var totalHotels int = 0
	if len(countStruct) > 0 {
		totalHotels = countStruct[0].Total
	}
	return &hotels, int(totalHotels), nil
}

func (r *hotelRepository) GetByUniqueKey(key string) (*models.Hotel, error) {
	var hotel models.Hotel
	stm := Hotel.SELECT(Hotel.AllColumns).
		FROM(Hotel.LEFT_JOIN(HotelRoom, Hotel.ID.EQ(HotelRoom.HotelID))).
		WHERE(Hotel.ID.EQ(String(key))).
		LIMIT(1)
	err := stm.Query(r.db, &hotel)

	if errors.Is(err, qrm.ErrNoRows) {
		return nil, hotelerrors.NotFoundError
	}
	if err != nil {
		return nil, err
	}
	return &hotel, nil
}

func (r *hotelRepository) SaveHotel(hotel *model.Hotel, rooms *[]model.HotelRoom) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	stmHotel := Hotel.INSERT(Hotel.AllColumns).
		MODEL(hotel)
	_, err = stmHotel.ExecContext(r.ctx, tx)
	if rooms != nil {
		stmRooms := HotelRoom.INSERT(HotelRoom.AllColumns).
			MODELS(rooms)
		_, err = stmRooms.ExecContext(r.ctx, tx)
	}
	if err != nil {
		return err
	}
	return nil
}

func (r *hotelRepository) UpdateHotel(hotel *model.Hotel, id string) error {
	err := r.gorm.Table("hotel").
		Where("id = ?", id).
		Updates(hotel).
		Error
	if err != nil {
		return err
	}
	return nil
}
