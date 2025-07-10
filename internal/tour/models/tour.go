package models

import "github.com/oliwallywonka/alpaca_backend/db/model"

type Tour struct {
	model.Tour
	Destinations []model.Destination
}