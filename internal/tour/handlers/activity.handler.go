package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/shared/utils"
	tadtos "github.com/oliwallywonka/alpaca_backend/internal/tour/dtos/touractivity"
)

func (h *TourHandler) CreateTourActivity(c echo.Context) error {
	resp := shared.NewAPI()
	dto := tadtos.CreateTourActivityDTO{}

	if err := c.Bind(&dto); err != nil {
		return resp.BindFailed(err)
	}

	if errors := utils.ValidateRequest(&dto); len(errors) > 0 {
		return resp.Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.ts.CreateTourActivity(c.Param("tourID"), &dto)
	if err != nil {
		return resp.Error(c, "CreateTourActivity", err, "error creating tour activity")
	}
	return c.JSON(resp.Created(nil))
}

func (h *TourHandler) UpdateTourActivity(c echo.Context) error {
	resp := shared.NewAPI()
	dto := tadtos.UpdateTourActivityDTO{}

	if err := c.Bind(&dto); err != nil {
		return resp.BindFailed(err)
	}

	if errors := utils.ValidateRequest(&dto); len(errors) > 0 {
		return resp.Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.ts.UpdateTourActivity(c.Param("tourActivityID"), &dto)
	if err != nil {
		return resp.Error(c, "UpdateTourActivity", err, "error updating tour activity")
	}
	return c.JSON(resp.Updated(nil))
}

func (h *TourHandler) GetTourActivities(c echo.Context) error {
	resp := shared.NewAPI()
	tourID := c.Param("tourID")

	activities, err := h.ts.GetTourActivities(tourID)
	if err != nil {
		return resp.Error(c, "GetTourActivities", err, "error getting tour activities")
	}
	return c.JSON(resp.Ok(activities))
}

func (h *TourHandler) DeleteTourActivity(c echo.Context) error {
	resp := shared.NewAPI()
	activityID := c.Param("tourActivityID")

	err := h.ts.DeleteTourActivity(activityID)
	if err != nil {
		return resp.Error(c, "DeleteTourActivity", err, "error deleting tour activity")
	}
	return c.JSON(resp.Deleted(nil))
}
