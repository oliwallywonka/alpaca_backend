package models

import (
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)


type Departure struct {
	ID             string        `db:"id"`
	TourID         string        `db:"tour_id"`
	State          string        `db:"state"`
	StartDate      int64         `db:"start_date"`
	EndDate        int64         `db:"end_date"`
	AvailableSlots int           `db:"available_slots"`
	CreatedAt      int64         `db:"created_at"`
	UpdatedAt      int64         `db:"updated_at"`
	Tour           *TourDeparture `db:"tour"`
}

type TourDeparture struct {
	Name       shared.LanguageField `db:"name"`
	RefPricePP float32              `db:"ref_price_pp"`
	Days       int8                 `db:"days"`
	Images     shared.Images        `db:"images"`
}
