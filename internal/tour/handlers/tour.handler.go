package handlers

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/shared/utils"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/services"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/tourerrors"
)

type TourHandler struct {
	ts       services.Service
	validate *validator.Validate
}

func NewTourHandler(ts services.Service) *TourHandler {
	return &TourHandler{
		ts:       ts,
		validate: validator.New(),
	}
}

func (h *TourHandler) NameExists(c echo.Context) error {
	dto := &dtos.ValidNameDTO{}
	res := shared.NewAPI()

	if err := c.Bind(dto); err != nil {
		return res.BindFailed(err)
	}

	if errors := utils.ValidateRequest(dto); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}

	exists, err := h.ts.NameOrSlugExists(dto.Name)
	if err != nil {
		return res.Error(c, "NameExists", err, "error checking tour name")
	}

	return c.JSON(res.Ok(exists))
}

func (h *TourHandler) GetTourByUniqueKey(c echo.Context) error {
	key := c.Param("key")
	res := shared.NewAPI()
	tour, err := h.ts.GetByUniqueKey(key)
	if err != nil {
		if errors.Is(err, tourerrors.NotFoundError) {
			return res.Error(c, "GetTourByUniqueKey", err, "tour not found")
		}
		return res.Error(c, "GetTourByUniqueKey", err, "error getting tour")
	}
	return c.JSON(res.Ok(tour))
}

func (h *TourHandler) SaveTour(c echo.Context) error {
	dto := &dtos.CreateTourDTO{}
	res := shared.NewAPI()

	if err := c.Bind(dto); err != nil {
		return shared.NewAPI().BindFailed(err)
	}

	if errors := utils.ValidateRequest(dto); len(errors) > 0 {
		return shared.NewAPI().Error(c, "ValidateRequest", nil, errors...)
	}

	tourID, err := h.ts.SaveTour(dto)
	if err != nil {
		if errors.Is(err, tourerrors.UniqueNameError) {
			return res.Error(c, "SaveTour", err, "tour name already exists")
		}
		if errors.Is(err, tourerrors.UniqueSlugError) {
			return res.Error(c, "SaveTour", err, "tour slug already exists")
		}
		return shared.NewAPI().Error(c, "SaveTour", err, "error creating tour")
	}
	fmt.Println(tourID)
	return c.JSON(res.Created(tourID))
}

func (h *TourHandler) GetPaginated(c echo.Context) error {
	dto := &shared.PaginatedQueryParamsDTO{}
	if err := echo.QueryParamsBinder(c).
		Int64("perPage", &dto.PerPage).
		Int64("page", &dto.Page).
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

	tours, err := h.ts.GetPaginated(dto)

	if err != nil {
		return shared.NewAPI().Error(c, "GetPaginated", err, "error getting paginated tour")
	}

	return c.JSON(shared.NewAPI().Ok(tours))
}

func (h *TourHandler) UpdateTour(c echo.Context) error {
	id := c.Param("id")
	dto := &dtos.UpdateTourDTO{}
	res := shared.NewAPI()

	if err := c.Bind(dto); err != nil {
		return shared.NewAPI().BindFailed(err)
	}

	if errors := utils.ValidateRequest(dto); len(errors) > 0 {
		return shared.NewAPI().Error(c, "ValidateRequest", nil, errors...)
	}
	err := h.ts.UpdateTour(dto, id)
	if err != nil {
		if errors.Is(err, tourerrors.NotFoundError) {
			return res.Error(c, "UpdateTour", err, "tour not found")
		}
		if errors.Is(err, tourerrors.UniqueNameError) {
			return res.Error(c, "UpdateTour", err, "tour name already exists")
		}
		if errors.Is(err, tourerrors.UniqueSlugError) {
			return res.Error(c, "UpdateTour", err, "tour slug already exists")
		}
		return shared.NewAPI().Error(c, "UpdateTour", err, "error updating tour")
	}
	return c.JSON(res.Updated(nil))
}
