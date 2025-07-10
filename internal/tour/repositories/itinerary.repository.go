package repositories

import (
	"errors"

	"github.com/go-jet/jet/v2/qrm"
	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	. "github.com/oliwallywonka/alpaca_backend/db/table"

	"github.com/oliwallywonka/alpaca_backend/internal/shared/errors"
)

func (r *tourRepository) GetItineraries(tourID string) (*[]model.Itinerary, error) {
	var itineraries []model.Itinerary
	stm := SELECT(Itinerary.AllColumns).
		FROM(Itinerary).
		WHERE(Itinerary.TourID.EQ(String(tourID))).
		ORDER_BY(Itinerary.Day.ASC())

	err := stm.Query(r.db, &itineraries)
	if err != nil {
		return nil, err
	}
	return &itineraries, nil
}

func (r *tourRepository) GetItineraryByID(id string) (*model.Itinerary, error) {
	var itinerary *model.Itinerary
	stm := SELECT(Itinerary.AllColumns).
		FROM(Itinerary).
		WHERE(Itinerary.ID.EQ(String(id))).
		LIMIT(1)
	err := stm.Query(r.db, &itinerary)

	if errors.Is(err, qrm.ErrNoRows) {
		return nil, commonerrors.NotFoundError
	}
	if err != nil {
		return nil, err
	}
	return itinerary, nil
}

func (r *tourRepository) CreateItinerary(itinerary *model.Itinerary) error {
	stm := Itinerary.INSERT(Itinerary.AllColumns).
		MODEL(itinerary)
	_, err := stm.Exec(r.db)
	if err != nil {
		return err
	}
	return nil
}

func (r *tourRepository) UpdateItinerary(itinerary *model.Itinerary, id string) error {
	return r.gorm.
		Table("itinerary").
		Where("id = ?", id).
		Updates(itinerary).
		Error
}

func (r *tourRepository) DeleteItinerary(tourID string) error {
	stm := Itinerary.DELETE().
		WHERE(Itinerary.ID.EQ(String(tourID)))
	_, err := stm.Exec(r.db)
	if err != nil {
		return err
	}
	return nil
}
