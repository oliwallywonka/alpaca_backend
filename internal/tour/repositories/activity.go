package repositories

import (
	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	. "github.com/oliwallywonka/alpaca_backend/db/table"

	"github.com/oliwallywonka/alpaca_backend/internal/tour/models"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/tourerrors"
)

func (r *tourRepository) GetTourActivities(tourID string) (*[]models.TourActivity, error) {
	var tourActivities []models.TourActivity
	stm := SELECT(TourActivity.AllColumns, Activity.AllColumns).
		FROM(
			TourActivity.LEFT_JOIN(TourActivity, TourActivity.ActivityID.EQ(Activity.ID)),
		).
		WHERE(TourActivity.TourID.EQ(String(tourID))).
		ORDER_BY(TourActivity.Day.DESC())

	err := stm.Query(r.db, &tourActivities)
	if err != nil {
		return nil, err
	}
	return &tourActivities, nil
}

func (r *tourRepository) GetTourActivityByID(activityID string) (*models.TourActivity, error) {
	var tourActivity models.TourActivity
	stm := SELECT(TourActivity.AllColumns, Activity.AllColumns).
		FROM(
			TourActivity.LEFT_JOIN(TourActivity, TourActivity.ActivityID.EQ(Activity.ID)),
		).
		WHERE(TourActivity.ID.EQ(String(activityID))).
		LIMIT(1)

	err := stm.Query(r.db, &tourActivity)
	if err != nil {
		return nil, err
	}

	if tourActivity.ID == "" {
		return nil, tourerrors.TourActivityNotFoundError
	}
	return &tourActivity, nil
}

func (r *tourRepository) CreateTourActivity(tourActivity *model.TourActivity) error {
	stm := TourActivity.INSERT(TourActivity.AllColumns).
		MODEL(tourActivity)
	_, err := stm.Exec(r.db)
	if err != nil {
		return err
	}
	return nil
}

func (r *tourRepository) UpdateTourActivity(tourActivityID string, tourActivity *model.TourActivity) error {
	return r.gorm.
		Table("tour_activity").
		Where("id = ?", tourActivityID).
		Updates(tourActivity).
		Error
}

func (r *tourRepository) DeleteTourActivity(tourActivityID string) error {
	smt := TourActivity.DELETE().
		WHERE(TourActivity.ID.EQ(String(tourActivityID)))
	_, err := smt.Exec(r.db)
	if err != nil {
		return err
	}
	return nil
}
