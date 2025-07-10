package repositories

import (
	"errors"
	"fmt"

	"github.com/go-jet/jet/v2/qrm"
	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	. "github.com/oliwallywonka/alpaca_backend/db/table"

	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	commonerrors "github.com/oliwallywonka/alpaca_backend/internal/shared/errors"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/models"
)

func (r *tourRepository) GetPaginated(params *shared.PaginatedQueryParams) (*[]models.Tour, int, error) {
	var tours *[]models.Tour
	var totalTours int64 = 0
	searchFilter := fmt.Sprintf("%%%s%%", params.SearchFilter)

	stm := SELECT(Tour.AllColumns, Destination.AllColumns).
		FROM(
			Tour.LEFT_JOIN(Tour, TourDestination.TourID.EQ(Tour.ID)).
				LEFT_JOIN(Destination, TourDestination.DestinationID.EQ(Destination.ID)),
		).
		WHERE(OR(
			CAST(Tour.Slug).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Tour.Name).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Tour.Accommodation).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Tour.Team).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Tour.Transport).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Tour.ShortDescription).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Tour.LongDescription).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Tour.RefPricePp).AS_TEXT().LIKE(String(searchFilter)),
		)).OFFSET(params.Offset).LIMIT(params.Limit)

	err := stm.Query(r.db, &tours)
	if err != nil {
		return nil, 0, err
	}

	stmCount := SELECT(COUNT(STAR).AS("total")).
		FROM(Tour).
		WHERE(OR(
			CAST(Tour.Slug).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Tour.Name).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Tour.Accommodation).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Tour.Team).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Tour.Transport).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Tour.ShortDescription).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Tour.LongDescription).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Tour.RefPricePp).AS_TEXT().LIKE(String(searchFilter)),
		))

	stmCount.Query(r.db, &totalTours)
	if err != nil {
		return nil, 0, err
	}
	return tours, int(totalTours), nil
}

func (r *tourRepository) GetByUniqueKey(key string) (*models.Tour, error) {
	var tour models.Tour
	stm := SELECT(Tour.AllColumns, Destination.AllColumns).
		FROM(
			Tour.LEFT_JOIN(Tour, TourDestination.TourID.EQ(Tour.ID)).
				LEFT_JOIN(Destination, TourDestination.DestinationID.EQ(Destination.ID)),
		).
		WHERE(Tour.ID.EQ(String(key))).
		LIMIT(1)

	err := stm.Query(r.db, &tour)

	if errors.Is(err, qrm.ErrNoRows) {
		return nil, commonerrors.NotFoundError
	}
	if err != nil {
		return nil, err
	}
	return &tour, nil
}

func (r *tourRepository) NameOrSlugExists(key string) (bool, error) {
	var totalTours int64 = 0
	stm := SELECT(COUNT(STAR).AS("total")).
		FROM(Tour).
		WHERE(OR(
			EXISTS(Raw(`
					SELECT 1
					FROM jsonb_each_text(name) AS name_key
					WHERE name_key.value = #1	
					`, RawArgs{"#1": key})),
			EXISTS(Raw(`
					SELECT 1
					FROM jsonb_each_text(slug) AS name_key
					WHERE name_key.value = #1	
					`, RawArgs{"#1": key})),
		))
	err := stm.Query(r.db, &totalTours)
	if err != nil {
		return false, err
	}
	return totalTours > 0, nil
}

func (r *tourRepository) SaveTour(tour *model.Tour) (*string, error) {
	stm := Destination.INSERT(Destination.AllColumns).
		MODEL(tour)
	_, err := stm.Exec(r.db)
	if err != nil {
		return nil, err
	}
	return &tour.ID, nil
}

func (r *tourRepository) UpdateTour(id string, tour *model.Tour) error {
	/* fmt.Printf("%+v", t)
	// Collect updates for JSON fields
	updateMap := map[string]interface{}{}

	// Update Name field if provided
	if t.Name != nil {
		updateMap["name"] = gorm.Expr("name || ?::jsonb", t.Name)
		t.Name = nil
	}

	// Update Slug field if provided
	if t.Slug != nil {
		updateMap["slug"] = gorm.Expr("slug || ?::jsonb", t.Slug)
		t.Slug = nil
	}

	// Update Transport field if provided
	if t.Transport != nil {
		updateMap["transport"] = gorm.Expr("transport || ?::jsonb", t.Transport)
		t.Transport = nil
	}

	// Update Accommodation field if provided
	if t.Accommodation != nil {
		updateMap["accommodation"] = gorm.Expr("accommodation || ?::jsonb", t.Accommodation)
		t.Accommodation = nil
	}

	// Update Team field if provided
	if t.Team != nil {
		updateMap["team"] = gorm.Expr("team || ?::jsonb", t.Team)
		t.Team = nil
	}

	// Update ShortDescription field if provided
	if t.ShortDescription != nil {
		updateMap["short_description"] = gorm.Expr("short_description || ?::jsonb", t.ShortDescription)
		t.ShortDescription = nil
	}

	// Update LongDescription field if provided
	if t.LongDescription != nil {
		updateMap["long_description"] = gorm.Expr("long_description || ?::jsonb", t.LongDescription)
		t.LongDescription = nil
	}

	// Apply partial updates for JSON fields if there are any
	if len(updateMap) > 0 {
		err := r.gorm.
			Table("tour").
			Where("id = ?", id).
			Updates(updateMap).
			Error
		if err != nil {
			return fmt.Errorf("exec: %w", err)
		}
	} */

	// Handle other non-JSON fields in a separate update
	err := r.gorm.
		Table("tour").
		Where("id = ?", id).
		Updates(tour).
		Error
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
