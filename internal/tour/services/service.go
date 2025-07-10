package services

import (
	"context"
	"mime/multipart"

	activityRepository "github.com/oliwallywonka/alpaca_backend/internal/activity/repositories"
	destRespository "github.com/oliwallywonka/alpaca_backend/internal/destination/repositories"
	mealRepository "github.com/oliwallywonka/alpaca_backend/internal/meal/repositories"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/dtos"
	tadtos "github.com/oliwallywonka/alpaca_backend/internal/tour/dtos/touractivity"
	tddtos "github.com/oliwallywonka/alpaca_backend/internal/tour/dtos/tourdestination"
	tidtos "github.com/oliwallywonka/alpaca_backend/internal/tour/dtos/touritinerary"
	tmdtos "github.com/oliwallywonka/alpaca_backend/internal/tour/dtos/tourmeal"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/repositories"
	cloudinaryfx "github.com/oliwallywonka/alpaca_backend/pkg/cloudinary"
	"github.com/oliwallywonka/alpaca_backend/settings"
)

type Service interface {
	SaveTour(tour *dtos.CreateTourDTO) (tourID *string, err error)
	GetTotalCount() int
	GetPaginated(dto *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error)
	GetByUniqueKey(slug string) (*dtos.TourDTO, error)
	UpdateTour(dto *dtos.UpdateTourDTO, id string) error
	NameOrSlugExists(name string) (bool, error)

	// IMAGES
	GetTourImages(id string) ([]string, error)
	SetImage(tourID string, imageFile multipart.File) error
	AddImage(tourID string, imageFile multipart.File) error
	DeleteImage(tourID string, imageID string) error

	// TOUR DESTINATIONS
	CreateTourDestination(dto *tddtos.CreateTourDestinationDTO, tourID string) error
	UpdateTourDestination(dto *tddtos.UpdateTourDestinationDTO, tourID string, tourDestinationID string) error
	GetTourDestinations(tourID string) (*[]tddtos.TourDestinationDTO, error)
	DeleteTourDestination(tourDestinationID string) error

	// ITINERARIES
	CreateItinerary(tourID string, itinerary *tidtos.CreateItineraryDTO) error
	UpdateItinerary(itineraryID string, dto *tidtos.UpdateItineraryDTO) error
	GetItineraries(tourID string) (*[]tidtos.ItineraryDTO, error)
	DeleteItinerary(itineraryID string) error
	SetItineraryImage(itineraryID string, imageFile multipart.File) error
	DeleteItineraryImage(itineraryID string, imageID string) error

	// MEALS
	GetTourMeals(tourID string) (*[]tmdtos.TourMealDTO, error)
	CreateTourMeal(tourID string, dto *tmdtos.CreateTourMealDTO) error
	UpdateTourMeal(tourMealID string, dto *tmdtos.UpdateTourMealDTO) error
	DeleteTourMeal(tourMealID string) error

	// ACTIVITIES
	CreateTourActivity(tourID string, dto *tadtos.CreateTourActivityDTO) error
	UpdateTourActivity(tourActivityID string, dto *tadtos.UpdateTourActivityDTO) error
	GetTourActivities(tourID string) (*[]tadtos.TourActivityDTO, error)
	DeleteTourActivity(tourActivityID string) error
}

type tourService struct {
	rep      repositories.Repository
	destRep  destRespository.Repository
	mealRep  mealRepository.Repository
	actRep   activityRepository.Repository
	clouds   cloudinaryfx.Service
	settings *settings.Settings
	ctx      context.Context
}

func New(
	rep repositories.Repository,
	destRep destRespository.Repository,
	mealRep mealRepository.Repository,
	actRep activityRepository.Repository,
	clouds cloudinaryfx.Service,
	s *settings.Settings,
	ctx context.Context,
) Service {
	return &tourService{
		rep:      rep,
		settings: s,
		ctx:      ctx,
		destRep:  destRep,
		clouds:   clouds,
		mealRep:  mealRep,
		actRep:   actRep,
	}
}
