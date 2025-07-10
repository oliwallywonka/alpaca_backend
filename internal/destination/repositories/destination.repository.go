package repositories

import (
	"fmt"

	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	. "github.com/oliwallywonka/alpaca_backend/db/table"

	"gorm.io/gorm/clause"

	"github.com/oliwallywonka/alpaca_backend/internal/destination/desterrors"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

func (r *destinationRepository) GetPaginated(params *shared.PaginatedQueryParams) (*[]model.Destination, int, error) {
	var countStruct []struct {
		Total int
	}

	searchFilter := fmt.Sprintf("%%%s%%", params.SearchFilter)
	stm := SELECT(Destination.AllColumns).
		FROM(Destination).
		WHERE(
			OR(
				CAST(Destination.Name).AS_TEXT().LIKE(String(searchFilter)),
				CAST(Destination.Description).AS_TEXT().LIKE(String(searchFilter)),
			),
		).
		LIMIT(params.Limit).
		OFFSET(params.Offset)
	var destinations []model.Destination
	err := stm.Query(r.db, &destinations)
	if err != nil {
		return nil, 0, err
	}

	stmCount := SELECT(COUNT(STAR).AS("total")).
		FROM(Destination).
		WHERE(
			OR(
				CAST(Destination.Name).AS_TEXT().LIKE(String(searchFilter)),
				CAST(Destination.Description).AS_TEXT().LIKE(String(searchFilter)),
			),
		)

	stmCount.Query(r.db, &countStruct)
	if err != nil {
		return nil, 0, err
	}
	var total int = 0
	if len(countStruct) > 0 {
		total = countStruct[0].Total
	}
	return &destinations, total, nil
}

func (r *destinationRepository) GetByUniqueKey(key string) (*[]model.Destination, error) {
	var destinations []model.Destination
	destCTE := CTE("cte")
	stm := WITH_RECURSIVE(
		destCTE.AS(
			SELECT(Destination.AllColumns).
				FROM(Destination).
				WHERE(Destination.ID.EQ(String(key))).
				UNION_ALL(
					SELECT(Destination.AllColumns).
						FROM(
							Destination.INNER_JOIN(
								destCTE,
								Destination.ID.From(destCTE).EQ(Destination.ParentID),
							),
						),
				),
		),
	)(
		SELECT(destCTE.AllColumns()).
			FROM(destCTE),
	)
	err := stm.Query(r.db, &destinations)
	if err != nil {
		return nil, desterrors.NotFoundError
	}
	return &destinations, err
}

func (r *destinationRepository) NameExists(name string) (bool, error) {
	var totalDestinations int64 = 0
	stm := SELECT(COUNT(Destination.Name)).
		FROM(Destination).
		WHERE(
			EXISTS(
				Raw(`
					SELECT 1
					FROM jsonb_each_text(name) AS name_key
					WHERE name_key.value = #1	
				`, RawArgs{"#1": name}),
			),
		)
	err := stm.Query(r.db, &totalDestinations)
	if err != nil {
		return false, err
	}
	return totalDestinations > 0, nil
}

func (r *destinationRepository) Create(destination *model.Destination) (*model.Destination, error) {
	var createdDestination model.Destination
	stm := Destination.INSERT(Destination.AllColumns).
		MODEL(destination).
		RETURNING(Destination.AllColumns)
	err := stm.Query(r.db, &createdDestination)

	if err != nil {
		return nil, err
	}
	return &createdDestination, nil
}

func (r *destinationRepository) Update(id string, destination *model.Destination) (*model.Destination, error) {
	err := r.gorm.
		Table("destination").
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(destination).
		Error
	if err != nil {
		return nil, err
	}
	return destination, nil
}
