package services

import (
	"mime/multipart"

	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/tourerrors"
)

func (s *tourService) GetTourImages(id string) ([]string, error) {
	images, err := s.rep.GetTourImages(id)
	if err != nil {
		return nil, err
	}
	return images, nil
}

func (s *tourService) SetImage(id string, imageFile multipart.File) error {
	tour, err := s.rep.GetByUniqueKey(id)
	if err != nil {
		return err
	}

	if tour.Images != nil || len(tour.Images) > 0 {
		for _, imageID := range tour.Images {
			_ = s.clouds.DeleteImage(imageID)
		}
	}
	uploadResult, err := s.clouds.UploadImage(imageFile)
	if err != nil {
		return err
	}
	imageID := []string{uploadResult.PublicID}

	err = s.rep.UpdateImageTour(id, imageID)
	if err != nil {
		return err
	}
	return nil
}

func (s *tourService) AddImage(tourID string, imageFile multipart.File) error {
	tour, err := s.rep.GetByUniqueKey(tourID)
	if err != nil {
		return err
	}
	uploadResult, err := s.clouds.UploadImage(imageFile)
	if err != nil {
		return err
	}
	images := tour.Images
	images = append(images, uploadResult.PublicID)
	err = s.rep.UpdateImageTour(tourID, images)
	if err != nil {
		return err
	}
	return nil
}

func (s *tourService) DeleteImage(tourID string, imageID string) error {
	hasImage := false
	imagePosition := 0

	tour, err := s.rep.GetByUniqueKey(tourID)
	if err != nil {
		return err
	}

	for index, image := range tour.Images {
		if image == imageID {
			hasImage = true
			imagePosition = index
		}
	}

	if !hasImage {
		return tourerrors.ImageNotFoundError
	}

	err = s.clouds.DeleteImage(imageID)
	if err != nil {
		return err
	}

	images := append(tour.Images[:imagePosition], tour.Images[imagePosition+1:]...)

	err = s.rep.UpdateTour(tourID, &model.Tour{Images: images})
	if err != nil {
		return err
	}
	return nil
}
