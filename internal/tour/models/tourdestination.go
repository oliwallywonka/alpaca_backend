package models

import (
	"github.com/oliwallywonka/alpaca_backend/db/model"
)

type TourDestination struct {
	model.TourDestination
	Destination model.Destination
}
