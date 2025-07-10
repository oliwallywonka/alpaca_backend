package userfx

import (
	"net/http"

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

func (h *Handler) GetUsers(c echo.Context) error {
	res := shared.NewAPI()
	dto := &shared.PaginatedQueryParamsDTO{
		OrderDirection: "DESC",
		OrderBy:        "created_at",
	}
	if err := echo.QueryParamsBinder(c).
		Int64("perPage", &dto.PerPage).
		Int64("page", &dto.Page).
		String("orderBy", &dto.OrderBy).
		String("orderDirection", &dto.OrderDirection).
		String("searchFilter", &dto.SearchFilter).
		BindError(); err != nil {
		return shared.NewAPI().BindFailed(err)
	}

	if errors := utils.ValidateRequest(dto); len(errors) > 0 {
		return shared.NewAPI().Error(c, "ValidateRequest", nil, errors...)
	}
	response, err := h.service.GetPaginated(dto)
	if err != nil {
		return res.Error(c, "GetPaginated", err, "error getting paginated users")
	}
	return c.JSON(res.Ok(response))
}

func (h *Handler) GetUser(c echo.Context) error {
	res := shared.NewAPI()
	userID := c.Param("userID")
	user, err := h.service.GetByUniqueKey(userID)
	if err != nil {
		return res.Error(c, "GetByUniqueKey", err, "error getting user")
	}
	return c.JSON(res.Ok(user))
}

func (h *Handler) Create(c echo.Context) error {
	res := shared.NewAPI()
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return res.BindFailed(err)
	}
	response, err := h.service.Save(user)
	if err != nil {
		return res.Error(c, "SaveUser", err, "error creating user")
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) Update(c echo.Context) error {
	res := shared.NewAPI()
	userID := c.Param("userID")
	var dto UpdateUserDTO
	if err := c.Bind(&dto); err != nil {
		return res.BindFailed(err)
	}
	user, err := h.service.Update(userID, &dto)
	if err != nil {
		return res.Error(c, "UpdateUser", err, "error updating user")
	}
	return c.JSON(http.StatusOK, user)
}
