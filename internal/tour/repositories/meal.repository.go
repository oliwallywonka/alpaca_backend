package repositories

import (
	"errors"

	"github.com/go-jet/jet/v2/qrm"
	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	. "github.com/oliwallywonka/alpaca_backend/db/table"

	"github.com/oliwallywonka/alpaca_backend/internal/tour/models"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/tourerrors"
)

func (r *tourRepository) GetTourMeals(tourID string) (*[]models.TourMeal, error) {
	var tourMeals *[]models.TourMeal
	stm := SELECT(TourMeals.AllColumns, Meal.AllColumns).
		FROM(
			TourMeals.LEFT_JOIN(TourMeals, TourMeals.MealID.EQ(Meal.ID)),
		).
		WHERE(TourMeals.TourID.EQ(String(tourID))).
		ORDER_BY(TourMeals.Day.ASC())

	err := stm.Query(r.db, &tourMeals)
	if err != nil {
		return nil, err
	}
	return tourMeals, nil
}

func (r *tourRepository) GetTourMealByID(id string) (*models.TourMeal, error) {
	var meal models.TourMeal
	stm := SELECT(TourMeals.AllColumns, Meal.AllColumns).
		FROM(
			TourMeals.LEFT_JOIN(TourMeals, TourMeals.MealID.EQ(Meal.ID)),
		).
		WHERE(TourMeals.ID.EQ(String(id))).
		LIMIT(1)
	err := stm.Query(r.db, &meal)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, tourerrors.TourMealNotFoundError
	}
	if err != nil {
		return nil, err
	}
	return &meal, nil
}

func (r *tourRepository) CreateTourMeal(meal *model.TourMeals) error {
	stm := TourMeals.INSERT(TourMeals.AllColumns).
		MODEL(meal)
	_, err := stm.Exec(r.db)
	if err != nil {
		return err
	}
	return nil
}

func (r *tourRepository) UpdateTourMealByDay(tourID string, day, newDay int32) error {
	return r.gorm.
		Table("tour_meals").
		Where("tour_id = ?", tourID).
		Where("day = ?", day).
		Update("day", newDay).
		Error
}

func (r *tourRepository) UpdateTourMeal(mealID string, meal *model.TourMeals) error {
	return r.gorm.
		Table("tour_meals").
		Where("id = ?", mealID).
		Updates(meal).
		Error
}

func (r *tourRepository) DeleteTourMeal(mealID string) error {
	stm := TourMeals.DELETE().
		WHERE(TourMeals.ID.EQ(String(mealID)))
	_, err := stm.Exec(r.db)
	if err != nil {
		return err
	}
	return nil
}
