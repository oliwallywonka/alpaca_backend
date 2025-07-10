package models

import (
	"github.com/oliwallywonka/alpaca_backend/db/model"
)

type Activity struct {
	model.Activity
	Destination model.Destination
}
