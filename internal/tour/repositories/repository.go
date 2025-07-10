package repositories

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/models"
)

type Repository interface {
	// TOURS
	GetPaginated(params *shared.PaginatedQueryParams) (*[]models.Tour, int, error)
	GetByUniqueKey(string) (*models.Tour, error)
	NameOrSlugExists(key string) (bool, error)
	SaveTour(tour *model.Tour) (tourID *string, err error)
	UpdateTour(id string, tour *model.Tour) error

	// IMAGES
	GetTourImages(tourID string) ([]string, error)
	UpdateImageTour(tourID string, images shared.ImageField) error

	// TOUR DESTINATIONS
	GetTourDestinations(tourID string) (*[]models.TourDestination, error)
	GetTourDestinationByID(tourDestinationID string) (*models.TourDestination, error)
	CreateTourDestination(tourDestination *model.TourDestination) error
	UpdateTourDestination(tourDestination *model.TourDestination, tourDestinationID string) error
	DeleteTourDestination(tourDestinationID string) error

	// ITINERARIES
	GetItineraries(tourID string) (*[]model.Itinerary, error)
	GetItineraryByID(itineraryID string) (*model.Itinerary, error)
	CreateItinerary(itinerary *model.Itinerary) error
	UpdateItinerary(itinerary *model.Itinerary, itineraryID string) error
	DeleteItinerary(itineraryID string) error

	// TOUR MEALS
	GetTourMeals(tourID string) (*[]models.TourMeal, error)
	GetTourMealByID(tourMealID string) (*models.TourMeal, error)
	CreateTourMeal(tourMeal *model.TourMeals) error
	UpdateTourMeal(mealID string, meal *model.TourMeals) error
	UpdateTourMealByDay(tourID string, day, newDay int32) error
	DeleteTourMeal(mealID string) error

	// TOUR ACTIVITIES
	GetTourActivities(tourID string) (*[]models.TourActivity, error)
	GetTourActivityByID(activityID string) (*models.TourActivity, error)
	CreateTourActivity(activity *model.TourActivity) error
	UpdateTourActivity(tourActivityID string, activity *model.TourActivity) error
	DeleteTourActivity(tourActivityID string) error
}

type tourRepository struct {
	db   *sql.DB
	gorm *gorm.DB
	ctx  context.Context
}

func New(db *sql.DB, gorm *gorm.DB, ctx context.Context) Repository {
	return &tourRepository{
		db:   db,
		gorm: gorm,
		ctx:  ctx,
	}
}
