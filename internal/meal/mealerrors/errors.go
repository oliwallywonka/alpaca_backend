package mealerrors

import "errors"

var (
	NotFoundError = errors.New("meal not found")
)