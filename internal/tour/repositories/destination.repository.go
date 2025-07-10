package repositories

import (
	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	. "github.com/oliwallywonka/alpaca_backend/db/table"

	"github.com/oliwallywonka/alpaca_backend/internal/tour/models"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/tourerrors"
)

func (r *tourRepository) GetTourDestinations(tourID string) (*[]models.TourDestination, error) {
	var tourDestinations []models.TourDestination
	stm := SELECT(TourDestination.AllColumns, Destination.AllColumns).
		FROM(
			TourDestination.LEFT_JOIN(TourDestination, TourDestination.DestinationID.EQ(Destination.ID)),
		).
		WHERE(TourDestination.TourID.EQ(String(tourID))).
		ORDER_BY(TourDestination.Day.DESC())

	err := stm.Query(r.db, &tourDestinations)
	if err != nil {
		return nil, err
	}
	return &tourDestinations, nil
}

func (r *tourRepository) GetTourDestinationByID(id string) (*models.TourDestination, error) {
	var destination *models.TourDestination
	stm := SELECT(TourDestination.AllColumns, Destination.AllColumns).
		FROM(
			TourDestination.LEFT_JOIN(TourDestination, TourDestination.DestinationID.EQ(Destination.ID)),
		).
		WHERE(TourDestination.ID.EQ(String(id))).
		LIMIT(1)
	err := stm.Query(r.db, &destination)
	if err != nil {
		return nil, err
	}

	if destination.ID == "" {
		return nil, tourerrors.TourDestNotFoundError
	}
	return destination, nil
}

func (r *tourRepository) CreateTourDestination(destination *model.TourDestination) error {
	stm := TourDestination.INSERT(TourDestination.AllColumns).
		MODEL(destination)
	_, err := stm.Exec(r.db)
	if err != nil {
		return err
	}
	return nil
}

func (r *tourRepository) UpdateTourDestination(destination *model.TourDestination, id string) error {
	return r.gorm.
		Table("tour_destination").
		Where("id = ?", id).
		Updates(destination).
		Error
}

func (r *tourRepository) DeleteTourDestination(tourDestinationID string) error {
	stm := TourDestination.DELETE().
		WHERE(TourDestination.ID.EQ(String(tourDestinationID)))
	_, err := stm.Exec(r.db)
	if err != nil {
		return err
	}
	return nil
}
