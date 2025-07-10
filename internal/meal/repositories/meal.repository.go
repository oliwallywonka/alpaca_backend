package repositories

import (
	"errors"
	"fmt"

	"github.com/go-jet/jet/v2/qrm"
	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	. "github.com/oliwallywonka/alpaca_backend/db/table"

	"github.com/oliwallywonka/alpaca_backend/internal/meal/mealerrors"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

func (r *mealRepository) GetPaginated(params *shared.PaginatedQueryParams) (*[]model.Meal, int, error) {
	var countStruct []struct {
		Total int
	}
	var meals *[]model.Meal
	searchFilter := fmt.Sprintf("%%%s%%", params.SearchFilter)
	stm := Meal.SELECT(Meal.AllColumns).
		FROM(Meal).
		WHERE(OR(
			CAST(Meal.Type).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Meal.Description).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Meal.RefPrice).AS_TEXT().LIKE(String(searchFilter)),
		)).
		LIMIT(params.Limit).
		OFFSET(params.Offset)
	err := stm.Query(r.db, &meals)
	if err != nil {
		return nil, 0, err
	}

	stmCount := SELECT(COUNT(STAR).AS("total")).
		FROM(Meal).
		WHERE(OR(
			CAST(Meal.Type).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Meal.Description).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Meal.RefPrice).AS_TEXT().LIKE(String(searchFilter)),
		))
	err = stmCount.Query(r.db, &countStruct)
	if err != nil {
		return nil, 0, err
	}
	var totalMeals int = 0
	if len(countStruct) > 0 {
		totalMeals = countStruct[0].Total
	}
	return meals, int(totalMeals), nil
}

func (r *mealRepository) GetByUniqueKey(key string) (*model.Meal, error) {
	var meal model.Meal
	stm := Meal.SELECT(Meal.AllColumns).
		FROM(Meal).
		WHERE(Meal.ID.EQ(String(key))).
		LIMIT(1)
	err := stm.Query(r.db, &meal)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, mealerrors.NotFoundError
	}
	if err != nil {
		return nil, err
	}
	return &meal, nil
}

func (r *mealRepository) Save(meal *model.Meal) error {
	stm := Meal.INSERT(Meal.AllColumns).
		MODEL(meal)
	_, err := stm.Exec(r.db)
	if err != nil {
		return err
	}
	return nil
}

func (r *mealRepository) Update(mealID string, meal *model.Meal) error {
	err := r.gorm.Table("meal").
		Where("id = ?", mealID).
		Updates(meal).
		Error
	if err != nil {
		return err
	}
	return nil
}
