package models

import "github.com/oliwallywonka/alpaca_backend/db/model"

// TODO add Contact[] jsonb COLUMN
type Hotel struct {
	model.Hotel
	Rooms []model.HotelRoom
}