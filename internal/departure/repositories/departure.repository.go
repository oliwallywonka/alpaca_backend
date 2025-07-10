package repositories

import (
	"fmt"

	"github.com/oliwallywonka/alpaca_backend/internal/departure/departureerrors"
	"github.com/oliwallywonka/alpaca_backend/internal/departure/models"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

func (r *departureRepository) GetDepartures(params *shared.PaginatedQueryParams) ([]*models.Departure, int, error) {
	var departures []*models.Departure = make([]*models.Departure, 0)
	var totalDepartures int64 = 0
	rows, err := r.db.
		Table("departure").
		Select(`
			departure.*, 
			tour.name, 
			tour.ref_price_pp, 
			tour.days, 
			tour.images
		`).
		Joins("LEFT JOIN tour ON tour.id = departure.tour_id").
		Where(`
			EXISTS (
				SELECT 1
				FROM jsonb_each_text(tour.name) as name_key
				WHERE name_key.value ILIKE CONCAT('%',?::text,'%')
			)
		`).
		Or("available_slots::text ILIKE CONCAT('%',?::text,'%')", params.SearchFilter).
		Group("departure.id").
		Order(fmt.Sprintf("%s %s", params.OrderBy, params.OrderDirection)).
		Offset(params.Offset).
		Limit(params.Limit).
		Rows()
	if err != nil {
		return nil, int(totalDepartures), err
	}

	defer rows.Close()

	for rows.Next() {
		var departure models.Departure
		err := rows.Scan(
			&departure.ID,
			&departure.TourID,
			&departure.State,
			&departure.StartDate,
			&departure.EndDate,
			&departure.AvailableSlots,
			&departure.CreatedAt,
			&departure.UpdatedAt,
			&departure.Tour.Name,
			&departure.Tour.RefPricePP,
			&departure.Tour.Days,
			&departure.Tour.Images,
		)
		if err != nil {
			return nil, int(totalDepartures), fmt.Errorf("query: %w", err)
		}
		departures = append(departures, &departure)
	}

	err = r.db.
		Table("departure").
		Joins("LEFT JOIN tour ON tour.id = departure.tour_id").
		Where(`
			EXISTS (
				SELECT 1
				FROM jsonb_each_text(tour.name) as name_key
				WHERE name_key.value ILIKE CONCAT('%',?::text,'%')
			)
		`).
		Or("available_slots::text ILIKE CONCAT('%',?::text,'%')", params.SearchFilter).
		Group("departure.id").
		Order(fmt.Sprintf("%s %s", params.OrderBy, params.OrderDirection)).
		Offset(params.Offset).
		Count(&totalDepartures).
		Error
	if err != nil {
		return nil, int(totalDepartures), fmt.Errorf("query: %w", err)
	}
	return departures, int(totalDepartures), nil
}

func (r *departureRepository) GetDepartureByUniqueKey(key string) (*models.Departure, error) {
	var departure models.Departure
	result := r.db.
		Select(`
			departure.*, 
			tour.name, 
			tour.ref_price_pp, 
			tour.days, 
			tour.images
		`).
		Where("id = ?", key).
		Limit(1).
		Table("departure").
		Find(&departure)

	if result.Error != nil {
		return nil, fmt.Errorf("query: %w", result.Error)
	}

	if departure.ID == "" {
		return nil, departureerrors.NotFoundError
	}
	return &departure, nil
}

func (r *departureRepository) SaveDeparture(departure *models.Departure) error {
	err := r.db.Table("departure").Create(departure).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *departureRepository) UpdateDeparture(departureID string, departure *models.Departure) error {
	err := r.db.
		Table("departure").
		Where("id = ?", departureID).
		Updates(departure).Error
	if err != nil {
		return err
	}
	return nil
}
