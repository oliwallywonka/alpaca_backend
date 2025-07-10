package repositories

import (
	"errors"
	"fmt"

	"github.com/go-jet/jet/v2/qrm"
	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	. "github.com/oliwallywonka/alpaca_backend/db/table"

	"github.com/oliwallywonka/alpaca_backend/internal/activity/activityerrors"
	"github.com/oliwallywonka/alpaca_backend/internal/activity/models"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

func (r *activityRepository) GetPaginated(params *shared.PaginatedQueryParams) (*[]models.Activity, int, error) {
	var countStruct []struct {
		Total int
	}
	var activities []models.Activity = make([]models.Activity, 0)

	searchFilter := fmt.Sprintf("%%%s%%", params.SearchFilter)
	stm := SELECT(Activity.AllColumns, Destination.AllColumns).
		FROM(Activity.LEFT_JOIN(Destination, Activity.DestinationID.EQ(Destination.ID))).
		WHERE(OR(
			CAST(Activity.Title).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Activity.Description).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Activity.RefPrice).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Destination.Name).AS_TEXT().LIKE(String(searchFilter)),
		)).
		LIMIT(params.Limit).
		OFFSET(params.Offset)
	err := stm.Query(r.db, &activities)
	if err != nil {
		return nil, 0, err
	}
	stmCount := SELECT(COUNT(STAR).AS("total")).
		FROM(Activity.LEFT_JOIN(Destination, Activity.DestinationID.EQ(Destination.ID))).
		WHERE(OR(
			CAST(Activity.Title).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Activity.Description).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Activity.RefPrice).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Destination.Name).AS_TEXT().LIKE(String(searchFilter)),
		))
	err = stmCount.Query(r.db, &countStruct)
	if err != nil {
		return nil, 0, err
	}

	var total int = 0
	if len(countStruct) > 0 {
		total = countStruct[0].Total
	}

	return &activities, total, nil
}

func (r *activityRepository) GetByUniqueKey(key string) (*models.Activity, error) {
	var activity models.Activity
	stm := Activity.SELECT(Activity.AllColumns).
		FROM(Activity.LEFT_JOIN(Destination, Activity.DestinationID.EQ(Destination.ID))).
		WHERE(Activity.ID.EQ(String(key))).
		LIMIT(1)
	err := stm.Query(r.db, &activity)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, activityerrors.NotFoundError
	}
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

func (r *activityRepository) Create(activity *model.Activity) error {
	fmt.Println(activity)
	stm := Activity.INSERT(Activity.AllColumns).
		MODEL(activity)
	_, err := stm.Exec(r.db)
	if err != nil {
		return nil
	}
	return nil
}

func (r *activityRepository) Update(activity *model.Activity, id string) error {
	err := r.gorm.
		Table("activity").
		Where("id = ?", id).
		Updates(activity).
		Error
	if err != nil {
		return err
	}
	return nil
}
