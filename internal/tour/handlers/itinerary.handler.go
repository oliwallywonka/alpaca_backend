package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/shared/utils"
	tidtos "github.com/oliwallywonka/alpaca_backend/internal/tour/dtos/touritinerary"
)

func (h *TourHandler) CreateItinerary(c echo.Context) error {
	res := shared.NewAPI()
	itinerary := tidtos.CreateItineraryDTO{}

	if err := c.Bind(&itinerary); err != nil {
		return res.BindFailed(err)
	}

	if errors := utils.ValidateRequest(&itinerary); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.ts.CreateItinerary(c.Param("tourID"), &itinerary)
	if err != nil {
		return res.Error(c, "CreateItinerary", err, "error creating itinerary")
	}
	return c.JSON(res.Created(nil))
}

func (h *TourHandler) UpdateItinerary(c echo.Context) error {
	res := shared.NewAPI()
	itineraryID := c.Param("itineraryID")
	dto := tidtos.UpdateItineraryDTO{}

	if err := c.Bind(&dto); err != nil {
		return res.BindFailed(err)
	}

	if errors := utils.ValidateRequest(&dto); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.ts.UpdateItinerary(itineraryID, &dto)
	if err != nil {
		return res.Error(c, "UpdateItinerary", err, "error updating itinerary")
	}
	return c.JSON(res.Updated(nil))
}

func (h *TourHandler) GetItineraries(c echo.Context) error {
	res := shared.NewAPI()
	tourID := c.Param("tourID")

	itineraries, err := h.ts.GetItineraries(tourID)
	if err != nil {
		return res.Error(c, "GetItineraries", err, "error getting itineraries")
	}
	return c.JSON(res.Ok(itineraries))
}

func (h *TourHandler) DeleteItinerary(c echo.Context) error {
	res := shared.NewAPI()

	itineraryID := c.Param("itineraryID")
	err := h.ts.DeleteItinerary(itineraryID)
	if err != nil {
		return res.Error(c, "DeleteItinerary", err, "error deleting itinerary")
	}
	return c.JSON(res.Deleted(nil))
}

func (h *TourHandler) SetItineraryImage(c echo.Context) error {
	res := shared.NewAPI()

	itineraryID := c.Param("itineraryID")
	file, err := c.FormFile("image")

	if err != nil {
		return res.Error(c, "CreateImage", err, "error getting image")
	}

	src, err := file.Open()
	if err != nil {
		return res.Error(c, "CreateImage", err, "error opening image")
	}
	defer src.Close()

	err = h.ts.SetItineraryImage(itineraryID, src)
	if err != nil {
		return res.Error(c, "CreateImage", err, "error uploading image to cloudinary")
	}
	return c.JSON(res.Created(nil))
}

func (h *TourHandler) DeleteItineraryImage(c echo.Context) error {
	res := shared.NewAPI()

	itineraryID := c.Param("itineraryID")
	imageID := c.Param("imageID")

	err := h.ts.DeleteItineraryImage(itineraryID, imageID)
	if err != nil {
		return res.Error(c, "DeleteImage", err, "error deleting image from cloudinary")
	}
	return c.JSON(res.Deleted(nil))
}
