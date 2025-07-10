package models

import (
	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/activity/models"
)

type TourActivity struct {
	model.TourActivity
	Activity models.Activity
}
