package handlers

import (
	"errors"

	"github.com/labstack/echo/v4"

	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/tourerrors"
)

func (h *TourHandler) GetTourImages(c echo.Context) error {
	id := c.Param("id")
	res := shared.NewAPI()
	images, err := h.ts.GetTourImages(id)
	if err != nil {
		if errors.Is(err, tourerrors.NotFoundError) {
			return res.Error(c, "GetTourImages", err, "tour not found")
		}
		return shared.NewAPI().Error(c, "GetTourImages", err, "error getting tour images")
	}
	return c.JSON(res.Ok(images))
}

func (h *TourHandler) SetImage(c echo.Context) error {
	res := shared.NewAPI()

	file, err := c.FormFile("image")
	id := c.Param("id")

	if err != nil {
		return res.Error(c, "CreateImage", err, "error getting image")
	}

	src, err := file.Open()
	if err != nil {
		return res.Error(c, "CreateImage", err, "error opening image")
	}
	defer src.Close()

	err = h.ts.SetImage(id, src)
	if err != nil {
		return res.Error(c, "CreateImage", err, "error uploading image to cloudinary")
	}
	return c.JSON(res.Created(nil))
}

func (h *TourHandler) AddImage(c echo.Context) error {
	res := shared.NewAPI()

	id := c.Param("id")
	file, err := c.FormFile("image")

	if err != nil {
		return res.Error(c, "CreateImage", err, "error getting image")
	}

	src, err := file.Open()
	if err != nil {
		return res.Error(c, "CreateImage", err, "error opening image")
	}
	defer src.Close()

	err = h.ts.AddImage(id, src)
	if err != nil {
		return res.Error(c, "CreateImage", err, "error uploading image to cloudinary")
	}
	return c.JSON(res.Ok(nil))
}

func (h *TourHandler) DeleteImage(c echo.Context) error {
	res := shared.NewAPI()

	id := c.Param("id")
	imageID := c.Param("imageID")

	err := h.ts.DeleteImage(id, imageID)
	if err != nil {
		if errors.Is(err, tourerrors.NotFoundError) {
			return res.Error(c, "DeleteImage", err, "tour not found")
		}
		if errors.Is(err, tourerrors.ImageNotFoundError) {
			return res.Error(c, "DeleteImage", err, "image not found")
		}
		return res.Error(c, "CreateImage", err, "error deleting image from cloudinary")
	}
	return c.JSON(res.Deleted(nil))
}
