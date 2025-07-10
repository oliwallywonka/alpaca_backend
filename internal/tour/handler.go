package tourfx

import (
	"github.com/labstack/echo/v4"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/shared/utils"
)

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetAll(c echo.Context) error {
	res := shared.NewAPI()
	dto := &shared.PaginatedQueryParamsDTO{}
	if err := echo.QueryParamsBinder(c).
		Int64("perPage", &dto.PerPage).
		Int64("page", &dto.Page).
		String("searchFilter", &dto.SearchFilter).
		String("orderBy", &dto.OrderBy).
		String("orderDirection", &dto.OrderDirection).
		BindError(); err != nil {
		return res.BindFailed(err)
	}
	if errors := utils.ValidateRequest(dto); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}
	response, err := h.service.GetAll(dto)
	if err != nil {
		return res.Error(c, "GetAll", err)
	}
	return c.JSON(res.Ok(response))
}

func (h *Handler) Save(c echo.Context) error {
	res := shared.NewAPI()
	var dto CreateTourDTO
	if err := c.Bind(&dto); err != nil {
		return res.BindFailed(err)
	}
	if errors := utils.ValidateRequest(&dto); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}
	response, err := h.service.Save(&dto)
	if err != nil {
		return res.Error(c, "SaveDestination", err)
	}
	return c.JSON(res.Ok(response))
}
