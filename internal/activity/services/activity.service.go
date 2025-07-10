package services

import (
	"mime/multipart"

	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/activity/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

func (s *activityService) GetPaginated(dto *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error) {
	activities, totalActivities, err := s.rep.GetPaginated(dto.DTOToModel())
	if err != nil {
		return nil, err
	}
	return &shared.PaginatedResponse{
		Items: dtos.ActivityDTOsModelToDTO(activities),
		Total: totalActivities,
		Page:  int(dto.Page),
	}, nil
}

func (s *activityService) GetByUniqueKey(key string) (*dtos.ActivityDTO, error) {
	activity, err := s.rep.GetByUniqueKey(key)
	if err != nil {
		return nil, err
	}
	return dtos.ActivityModelToDTO(activity), nil
}

func (s *activityService) Save(activity *dtos.CreateActivityDTO) error {
	err := s.rep.Create(activity.ToModel())
	if err != nil {
		return err
	}
	return nil
}

func (s *activityService) Update(dto *dtos.UpdateActivityDTO, id string) error {
	err := s.rep.Update(dto.DTOToModel(), id)
	if err != nil {
		return err
	}
	return nil
}

func (s *activityService) SetImage(id string, imageFile multipart.File) error {
	activity, err := s.rep.GetByUniqueKey(id)
	if err != nil {
		return err
	}

	if len(activity.Images) > 0 {
		for _, imageID := range activity.Images {
			_ = s.clods.DeleteImage(imageID)
		}
	}

	uploadResult, err := s.clods.UploadImage(imageFile)
	if err != nil {
		return err
	}
	imageID := []string{uploadResult.PublicID}

	err = s.rep.Update(&model.Activity{Images: imageID}, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *activityService) DeleteImage(activityID string, imageID string) error {
	hasImage := false
	imagePosition := 0

	activity, err := s.rep.GetByUniqueKey(activityID)
	if err != nil {
		return err
	}

	for index, image := range activity.Images {
		if image == imageID {
			hasImage = true
			imagePosition = index
		}
	}

	if !hasImage {
		return nil
	}

	err = s.clods.DeleteImage(imageID)
	if err != nil {
		return err
	}

	images := append(activity.Images[:imagePosition], activity.Images[imagePosition+1:]...)

	err = s.rep.Update(&model.Activity{Images: images}, activityID)
	if err != nil {
		return err
	}
	return nil
}
