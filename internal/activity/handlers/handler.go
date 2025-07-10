package handlers

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/oliwallywonka/alpaca_backend/internal/activity/activityerrors"
	"github.com/oliwallywonka/alpaca_backend/internal/activity/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/activity/services"
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

func (h *Handler) GetActivityByUniqueKey(c echo.Context) error {
	res := shared.NewAPI()

	key := c.Param("key")

	activity, err := h.s.GetByUniqueKey(key)

	if errors.Is(err, activityerrors.NotFoundError) {
		return res.Error(c, "GetActivityByUniqueKey", err, "activity not found")
	}

	if err != nil {
		return res.Error(c, "GetActivityByUniqueKey", err, "error getting activity")
	}
	return c.JSON(res.Ok(activity))
}

func (h *Handler) Save(c echo.Context) error {
	res := shared.NewAPI()
	dto := &dtos.CreateActivityDTO{}

	if err := c.Bind(dto); err != nil {
		return res.BindFailed(err)
	}

	if errors := utils.ValidateRequest(dto); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.s.Save(dto)
	if err != nil {
		return res.Error(c, "SaveActivity", err, "error creating activity")
	}
	return c.JSON(res.Created(nil))
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

	activities, err := h.s.GetPaginated(dto)

	if err != nil {
		return res.Error(c, "GetPaginated", err, "error getting paginated activity")
	}

	return c.JSON(res.Ok(activities))
}

func (h *Handler) Update(c echo.Context) error {
	res := shared.NewAPI()
	id := c.Param("id")
	dto := &dtos.UpdateActivityDTO{}
	if err := c.Bind(dto); err != nil {
		return res.BindFailed(err)
	}

	if errors := utils.ValidateRequest(dto); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.s.Update(dto, id)
	if err != nil {
		return res.Error(c, "UpdateActivity", err, "error updating activity")
	}
	return c.JSON(res.Updated(nil))
}

func (h *Handler) SetImage(c echo.Context) error {
	res := shared.NewAPI()

	activityID := c.Param("activityID")
	file, err := c.FormFile("image")

	if err != nil {
		return res.Error(c, "CreateImage", err, "error getting image")
	}

	src, err := file.Open()
	if err != nil {
		return res.Error(c, "CreateImage", err, "error opening image")
	}
	defer src.Close()

	err = h.s.SetImage(activityID, src)
	if err != nil {
		return res.Error(c, "CreateImage", err, "error uploading image to cloudinary")
	}
	return c.JSON(res.Ok(nil))
}

func (h *Handler) DeleteImage(c echo.Context) error {
	res := shared.NewAPI()

	activityID := c.Param("activityID")
	imageID := c.Param("imageID")

	err := h.s.DeleteImage(activityID, imageID)
	if err != nil {
		return res.Error(c, "DeleteImage", err, "error deleting image from cloudinary")
	}
	return c.JSON(res.Ok(nil))
}
