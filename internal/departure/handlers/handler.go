package handler

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/oliwallywonka/alpaca_backend/internal/departure/departureerrors"
	"github.com/oliwallywonka/alpaca_backend/internal/departure/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/departure/services"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/shared/utils"
)

type DepartureHandler struct {
	ds services.Service
}

func NewDepartureHandler(ds services.Service) *DepartureHandler {
	return &DepartureHandler{
		ds: ds,
	}
}

func (h *DepartureHandler) GetDepartures(c echo.Context) error {
	dto := &shared.PaginatedQueryParamsDTO{}
	if err := echo.QueryParamsBinder(c).
		Int("perPage", &dto.PerPage).
		Int("page", &dto.Page).
		String("orderBy", &dto.OrderBy).
		String("orderDirection", &dto.OrderDirection).
		String("searchFilter", &dto.SearchFilter).
		BindError(); err != nil {
		return shared.NewAPI().BindFailed(err)
	}
	dto.DefaultValues()

	if errors := utils.ValidateRequest(dto); len(errors) > 0 {
		return shared.NewAPI().Error(c, "ValidateRequest", nil, errors...)
	}

	departures, err := h.ds.GetDepartures(dto)
	if err != nil {
		return shared.NewAPI().Error(c, "GetDepartures", err, "error getting paginated departures")
	}
	return c.JSON(shared.NewAPI().Ok(departures))
}

func (h *DepartureHandler) GetDepartureByUniqueKey(c echo.Context) error {
	key := c.Param("key")
	res := shared.NewAPI()
	departure, err := h.ds.GetDepartureByUniqueKey(key)
	if err != nil {
		if errors.Is(err, departureerrors.NotFoundError) {
			return res.Error(c, "GetDepartureByUniqueKey", err, "departure not found")
		}
		return res.Error(c, "GetDepartureByUniqueKey", err, "error getting departure")
	}
	return c.JSON(res.Ok(departure))
}

func (h *DepartureHandler) SaveDeparture(c echo.Context) error {
	dto := &dtos.CreateDepartureDTO{}
	res := shared.NewAPI()

	if err := c.Bind(dto); err != nil {
		return shared.NewAPI().BindFailed(err)
	}

	if errors := utils.ValidateRequest(dto); len(errors) > 0 {
		return shared.NewAPI().Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.ds.SaveDeparture(dto)
	if err != nil {
		return shared.NewAPI().Error(c, "SaveDeparture", err, "error creating departure")
	}
	return c.JSON(res.Created(nil))
}

func (h *DepartureHandler) UpdateDeparture(c echo.Context) error {
	id := c.Param("id")
	dto := &dtos.UpdateDepartureDTO{}
	res := shared.NewAPI()

	if err := c.Bind(dto); err != nil {
		return shared.NewAPI().BindFailed(err)
	}

	if errors := utils.ValidateRequest(dto); len(errors) > 0 {
		return shared.NewAPI().Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.ds.UpdateDeparture(id, dto)
	if err != nil {
		if errors.Is(err, departureerrors.NotFoundError) {
			return res.Error(c, "UpdateDeparture", err, "departure not found")
		}
		return shared.NewAPI().Error(c, "UpdateDeparture", err, "error updating departure")
	}
	return c.JSON(res.Updated(nil))
}
