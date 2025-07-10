package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/shared/utils"
	tmdtos "github.com/oliwallywonka/alpaca_backend/internal/tour/dtos/tourmeal"
)

func (h *TourHandler) CreateTourMeal(c echo.Context) error {
	resp := shared.NewAPI()
	dto := tmdtos.CreateTourMealDTO{}

	if err := c.Bind(&dto); err != nil {
		return resp.BindFailed(err)
	}
	fmt.Printf("%+v\n", c.Param("tourID"))
	if errors := utils.ValidateRequest(&dto); len(errors) > 0 {
		return resp.Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.ts.CreateTourMeal(c.Param("tourID"), &dto)
	if err != nil {
		return resp.Error(c, "CreateTourMeal", err, "error creating tour meal")
	}
	return c.JSON(resp.Created(nil))
}

func (h *TourHandler) UpdateTourMeal(c echo.Context) error {
	resp := shared.NewAPI()
	dto := tmdtos.UpdateTourMealDTO{}

	if err := c.Bind(&dto); err != nil {
		return resp.BindFailed(err)
	}

	if errors := utils.ValidateRequest(&dto); len(errors) > 0 {
		return resp.Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.ts.UpdateTourMeal(c.Param("tourMealID"), &dto)
	if err != nil {
		return resp.Error(c, "UpdateTourMeal", err, "error updating tour meal")
	}
	return c.JSON(resp.Ok(nil))
}

func (h *TourHandler) GetTourMeals(c echo.Context) error {
	resp := shared.NewAPI()
	tourID := c.Param("tourID")

	meals, err := h.ts.GetTourMeals(tourID)
	if err != nil {
		return resp.Error(c, "GetTourMeals", err, "error getting tour meals")
	}
	return c.JSON(resp.Ok(meals))
}

func (h *TourHandler) DeleteTourMeal(c echo.Context) error {
	resp := shared.NewAPI()
	mealID := c.Param("tourMealID")

	err := h.ts.DeleteTourMeal(mealID)
	if err != nil {
		return resp.Error(c, "DeleteTourMeal", err, "error deleting tour meal")
	}
	return c.JSON(resp.Deleted(nil))
}
