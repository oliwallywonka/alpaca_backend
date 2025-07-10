package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/oliwallywonka/alpaca_backend/internal/hotel/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/hotel/services"
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
	dto.DefaultValues()

	if errors := utils.ValidateRequest(dto); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}

	hotels, err := h.s.GetPaginated(dto)

	if err != nil {
		return res.Error(c, "GetPaginated", err, "error getting paginated hotel")
	}

	return c.JSON(res.Ok(hotels))
}

func (h *Handler) GetByUniqueKey(c echo.Context) error {
	res := shared.NewAPI()
	key := c.Param("key")

	hotel, err := h.s.GetByUniqueKey(key)
	if err != nil {
		return res.Error(c, "GetByUniqueKey", err, "error getting hotel by unique key")
	}
	return c.JSON(res.Ok(hotel))
}

func (h *Handler) SaveHotel(c echo.Context) error {
	res := shared.NewAPI()
	hotel := &dtos.CreateHotelDTO{}

	if err := c.Bind(hotel); err != nil {
		return res.BindFailed(err)
	}

	if errors := utils.ValidateRequest(hotel); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.s.SaveHotel(hotel)
	if err != nil {
		return res.Error(c, "SaveHotel", err, "error creating hotel")
	}
	return c.JSON(res.Created(nil))
}

func (h *Handler) UpdateHotel(c echo.Context) error {
	res := shared.NewAPI()
	id := c.Param("id")
	hotel := &dtos.UpdateHotelDTO{}

	if err := c.Bind(hotel); err != nil {
		return res.BindFailed(err)
	}

	if errors := utils.ValidateRequest(hotel); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.s.UpdateHotel(hotel, id)
	if err != nil {
		return res.Error(c, "UpdateHotel", err, "error updating hotel")
	}
	return c.JSON(res.Updated(nil))
}
