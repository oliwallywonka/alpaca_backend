package services

import (
	"mime/multipart"

	"github.com/oliwallywonka/alpaca_backend/internal/activity/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/activity/repositories"
	"github.com/oliwallywonka/alpaca_backend/pkg/cloudinary"
)

type Service interface {
	GetPaginated(dto *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error)
	GetByUniqueKey(key string) (*dtos.ActivityDTO, error)
	Save(activity *dtos.CreateActivityDTO) error
	Update(dto *dtos.UpdateActivityDTO, id string) error

	// IMAGES
	SetImage(activityID string, file multipart.File) error
	DeleteImage(activityID string, imageID string) error
}

type activityService struct {
	rep   repositories.Repository
	clods cloudinaryfx.Service
}

func New(rep repositories.Repository, clods cloudinaryfx.Service) Service {
	return &activityService{
		rep:   rep,
		clods: clods,
	}
}
