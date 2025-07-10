package destinationfx

import (
	"github.com/labstack/echo/v4"
	"github.com/oliwallywonka/alpaca_backend/db/model"
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
	dto := &shared.PaginatedQueryParamsDTO{
		OrderDirection: "DESC",
		OrderBy:        "destination.created_at",
	}
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
	response, err := h.service.GetAll(dto)
	if err != nil {
		return res.Error(c, "GetAllDestinations", err)
	}
	return c.JSON(res.Ok(response))
}

func (h *Handler) GetByID(c echo.Context) error {
	res := shared.NewAPI()
	destinationID := c.Param("destinationID")
	response, err := h.service.GetByID(destinationID)
	if err != nil {
		return res.Error(c, "GetDestinationByID", err)
	}
	return c.JSON(res.Ok(response))
}

func (h *Handler) Create(c echo.Context) error {
	res := shared.NewAPI()
	var dto model.Destination
	if err := c.Bind(&dto); err != nil {
		return res.BindFailed(err)
	}
	if errors := utils.ValidateRequest(&dto); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}
	response, err := h.service.Save(&dto)
	if err != nil {
		return res.Error(c, "CreateDestination", err)
	}
	return c.JSON(res.Ok(response))
}

func (h *Handler) Update(c echo.Context) error {
	res := shared.NewAPI()
	destinationID := c.Param("destinationID")
	var dto model.Destination
	if err := c.Bind(&dto); err != nil {
		return res.BindFailed(err)
	}
	if errors := utils.ValidateRequest(&dto); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}
	response, err := h.service.Update(destinationID, &dto)
	if err != nil {
		return res.Error(c, "UpdateDestination", err)
	}
	return c.JSON(res.Ok(response))
}
