package rolefx

import (
	"github.com/labstack/echo/v4"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/shared/utils"
)

type Handler struct {
	s IService
}

func NewHandler(s IService) *Handler {
	return &Handler{
		s: s,
	}
}

func (h *Handler) GetRoles(c echo.Context) error {
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

	roles, err := h.s.GetPaginated(dto)

	if err != nil {
		return res.Error(c, "GetPaginated", err, "error getting paginated roles")
	}

	return c.JSON(res.Ok(roles))
}

func (h *Handler) GetRole(c echo.Context) error {
	res := shared.NewAPI()
	roleID := c.Param("roleID")
	role, err := h.s.GetByUniqueKey(roleID)
	if err != nil {
		return res.Error(c, roleID, err)
	}
	return c.JSON(res.Ok(role))
}

func (h *Handler) Create(c echo.Context) error {
	res := shared.NewAPI()
	role := &model.Role{}
	if err := c.Bind(role); err != nil {
		return res.BindFailed(err)
	}

	role, err := h.s.Save(role)
	if err != nil {
		return res.Error(c, "Save", err)
	}
	return c.JSON(res.Ok(role))
}

func (h *Handler) Update(c echo.Context) error {
	res := shared.NewAPI()
	roleID := c.Param("roleID")
	dto := &UpdateRoleDTO{}
	if err := c.Bind(dto); err != nil {
		return res.BindFailed(err)
	}

	role, err := h.s.Update(roleID, dto)
	if err != nil {
		return res.Error(c, "Update", err)
	}
	return c.JSON(res.Ok(role))
}
