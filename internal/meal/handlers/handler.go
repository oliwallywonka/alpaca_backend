package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/oliwallywonka/alpaca_backend/internal/meal/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/meal/services"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/shared/utils"
)

type Handler struct {
	s services.Service
}

func New(s services.Service) *Handler {
	return &Handler{
		s: s,
	}
}

func (h *Handler) SaveMeal(c echo.Context) error {
	res := shared.NewAPI()
	meal := &dtos.CreateMealDTO{}

	if err := c.Bind(meal); err != nil {
		return res.BindFailed(err)
	}

	if errors := utils.ValidateRequest(meal); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.s.Save(meal)
	if err != nil {
		return res.Error(c, "SaveMeal", err, "error creating meal")
	}
	return c.JSON(res.Created(nil))
}

func (h *Handler) UpdateMeal(c echo.Context) error {
	res := shared.NewAPI()
	mealID := c.Param("mealID")
	meal := &dtos.UpdateMealDTO{}

	if err := c.Bind(meal); err != nil {
		return res.BindFailed(err)
	}

	if errors := utils.ValidateRequest(meal); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.s.Update(mealID, meal)
	if err != nil {
		return res.Error(c, "UpdateMeal", err, "error updating meal")
	}
	return c.JSON(res.Updated(nil))
}

func (h *Handler) GetPaginated(c echo.Context) error {
	res := shared.NewAPI()
	dto := &shared.PaginatedQueryParamsDTO{}
	if err := echo.QueryParamsBinder(c).
		Int64("perPage", &dto.PerPage).
		Int64("page", &dto.Page).
		String("orderBy", &dto.OrderBy).
		String("orderDirection", &dto.OrderDirection).
		String("searchFilter", &dto.SearchFilter).
		BindError(); err != nil {
		return res.BindFailed(err)
	}

	if errors := utils.ValidateRequest(dto); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}

	meals, err := h.s.GetPaginated(dto)

	if err != nil {
		return res.Error(c, "GetPaginated", err, "error getting paginated meal")
	}

	return c.JSON(res.Ok(meals))
}

func (h *Handler) GetByUniqueKey(c echo.Context) error {
	res := shared.NewAPI()
	key := c.Param("mealID")

	meal, err := h.s.GetByUniqueKey(key)
	if err != nil {
		return res.Error(c, "GetByUniqueKey", err, "error getting meal")
	}
	return c.JSON(res.Ok(meal))
}
