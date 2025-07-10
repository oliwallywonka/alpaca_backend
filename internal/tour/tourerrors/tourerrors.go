package tourerrors

import "errors"

var (
	UniqueNameError           = errors.New("tour name already exists")
	UniqueSlugError           = errors.New("tour slug already exists")
	NotFoundError             = errors.New("tour not found")
	DataBaseError             = errors.New("database error")
	ImageNotFoundError        = errors.New("image not found")
	TourDestNotFoundError     = errors.New("tour destination not found")
	ItineraryNotFoundError    = errors.New("itinerary not found")
	TourMealNotFoundError     = errors.New("tour meal not found")
	TourActivityNotFoundError = errors.New("tour activity not found")
)
