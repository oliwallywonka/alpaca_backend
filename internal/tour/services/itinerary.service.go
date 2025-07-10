package services

import (
	"mime/multipart"

	"github.com/oliwallywonka/alpaca_backend/db/model"
	tidtos "github.com/oliwallywonka/alpaca_backend/internal/tour/dtos/touritinerary"
)

func (s *tourService) CreateItinerary(tourID string, itinerary *tidtos.CreateItineraryDTO) error {
	_, err := s.rep.GetByUniqueKey(tourID)
	if err != nil {
		return err
	}

	err = s.rep.CreateItinerary(itinerary.DTOToModel(tourID))
	if err != nil {
		return err
	}
	return nil
}

func (s *tourService) UpdateItinerary(itineraryID string, dto *tidtos.UpdateItineraryDTO) error {
	_, err := s.rep.GetItineraryByID(itineraryID)
	if err != nil {
		return err
	}

	itinerary := dto.DTOToModel()
	err = s.rep.UpdateItinerary(itinerary, itineraryID)
	if err != nil {
		return err
	}
	return nil
}

func (s *tourService) GetItineraries(tourID string) (*[]tidtos.ItineraryDTO, error) {
	_, err := s.rep.GetByUniqueKey(tourID)
	if err != nil {
		return nil, err
	}
	itineraries, err := s.rep.GetItineraries(tourID)
	if err != nil {
		return nil, err
	}
	dtos := tidtos.ItineraryDTOsModelToDTO(itineraries)
	return dtos, nil
}

func (s *tourService) DeleteItinerary(itineraryID string) error {
	itineraryDB, err := s.rep.GetItineraryByID(itineraryID)
	if err != nil {
		return err
	}
	err = s.rep.DeleteItinerary(itineraryID)
	if err != nil {
		return err
	}

	// DELETE ALL IMAGES
	for _, imageID := range itineraryDB.Images {
		err = s.clouds.DeleteImage(imageID)
		if err != nil {
			return err
		}
	}

	itineraries, err := s.rep.GetItineraries(itineraryDB.TourID)

	if err != nil {
		return err
	}

	/* ITS NECESARY REORDER ALL ITINERARIES DAY AFTER DELETE */
	if len(*itineraries) == 0 {
		return nil
	}

	for index, itinerary := range *itineraries {
		err = s.rep.UpdateItinerary(&model.Itinerary{Day: int32(index + 1)}, itinerary.ID)
		if err != nil {
			return err
		}

		err = s.rep.UpdateTourMealByDay(itinerary.TourID, itinerary.Day, int32(index+1))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *tourService) SetItineraryImage(itineraryID string, imageFile multipart.File) error {
	itinerary, err := s.rep.GetItineraryByID(itineraryID)
	if err != nil {
		return err
	}

	if len(itinerary.Images) > 0 {
		for _, imageID := range itinerary.Images {
			_ = s.clouds.DeleteImage(imageID)
		}
	}

	uploadResult, err := s.clouds.UploadImage(imageFile)
	if err != nil {
		return err
	}
	imageID := []string{uploadResult.PublicID}

	err = s.rep.UpdateItinerary(&model.Itinerary{Images: imageID}, itineraryID)
	if err != nil {
		return err
	}
	return nil
}

func (s *tourService) DeleteItineraryImage(itineraryID string, imageID string) error {
	hasImage := false
	imagePosition := 0

	itinerary, err := s.rep.GetItineraryByID(itineraryID)
	if err != nil {
		return err
	}

	for index, image := range itinerary.Images {
		if image == imageID {
			hasImage = true
			imagePosition = index
		}
	}

	if !hasImage {
		return nil
	}

	err = s.clouds.DeleteImage(imageID)
	if err != nil {
		return err
	}

	images := append(itinerary.Images[:imagePosition], itinerary.Images[imagePosition+1:]...)

	err = s.rep.UpdateItinerary(&model.Itinerary{Images: images}, itineraryID)
	if err != nil {
		return err
	}
	return nil
}
