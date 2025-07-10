package models

import (
	"github.com/oliwallywonka/alpaca_backend/db/model"
)

type TourMeal struct {
	model.TourMeals
	Meal model.Meal
}
