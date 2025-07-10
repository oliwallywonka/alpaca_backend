package handlers

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/shared/utils"
	tddtos "github.com/oliwallywonka/alpaca_backend/internal/tour/dtos/tourdestination"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/tourerrors"
)

func (h *TourHandler) CreateTourDestination(c echo.Context) error {
	resp := shared.NewAPI()
	dto := tddtos.CreateTourDestinationDTO{}

	if err := c.Bind(&dto); err != nil {
		return resp.BindFailed(err)
	}

	if errors := utils.ValidateRequest(&dto); len(errors) > 0 {
		return resp.Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.ts.CreateTourDestination(&dto, c.Param("tourID"))
	if err != nil {
		return resp.Error(c, "CreateTourDestination", err, "error creating tour destination")
	}
	return c.JSON(resp.Created(nil))
}

func (h *TourHandler) UpdateTourDestination(c echo.Context) error {
	resp := shared.NewAPI()
	dto := tddtos.UpdateTourDestinationDTO{}

	if err := c.Bind(&dto); err != nil {
		return resp.BindFailed(err)
	}

	if errors := utils.ValidateRequest(&dto); len(errors) > 0 {
		return resp.Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.ts.UpdateTourDestination(&dto, c.Param("tourID"), c.Param("tourDestinationID"))
	if err != nil {
		if errors.Is(err, tourerrors.NotFoundError) {
			return resp.Error(c, "UpdateTourDestination", err, "tour not found")
		}
		if errors.Is(err, tourerrors.TourDestNotFoundError) {
			return resp.Error(c, "UpdateTourDestination", err, "tour destination not found")
		}
		return resp.Error(c, "UpdateTourDestination", err, "error updating tour destination")
	}
	return c.JSON(resp.Updated(nil))
}

func (h *TourHandler) GetTourDestinations(c echo.Context) error {
	resp := shared.NewAPI()
	tourID := c.Param("tourID")

	destinations, err := h.ts.GetTourDestinations(tourID)
	if err != nil {
		return resp.Error(c, "GetTourDestinations", err, "error getting tour destinations")
	}
	return c.JSON(resp.Ok(destinations))
}

func (h *TourHandler) DeleteTourDestination(c echo.Context) error {
	resp := shared.NewAPI()
	destinationID := c.Param("tourDestinationID")

	err := h.ts.DeleteTourDestination(destinationID)
	if err != nil {
		return resp.Error(c, "DeleteTourDestination", err, "error deleting tour destination")
	}
	return c.JSON(resp.Deleted(nil))
}
