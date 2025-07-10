package providerfx

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

func (h *Handler) GetAll(c echo.Context) error {
	res := shared.NewAPI()
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

	if errors := utils.ValidateRequest(dto); len(errors) > 0 {
		return shared.NewAPI().Error(c, "ValidateRequest", nil, errors...)
	}
	response, err := h.service.GetPaginated(dto)
	if err != nil {
		return res.Error(c, "GetPaginated", err, "error getting paginated providers")
	}
	return c.JSON(res.Ok(response))
}

func (h *Handler) GetProvider(c echo.Context) error {
	res := shared.NewAPI()
	providerID := c.Param("providerID")
	provider, err := h.service.GetByUniqueKey(providerID)
	if err != nil {
		return res.Error(c, "GetByUniqueKey", err, "error getting provider")
	}
	return c.JSON(res.Ok(provider))
}

func (h *Handler) Create(c echo.Context) error {
	res := shared.NewAPI()
	provider := &model.Provider{}
	if err := c.Bind(provider); err != nil {
		return res.BindFailed(err)
	}
	response, err := h.service.Save(provider)
	if err != nil {
		return res.Error(c, "SaveProvider", err, "error creating provider")
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) Update(c echo.Context) error {
	res := shared.NewAPI()
	providerID := c.Param("providerID")
	var dto UpdateProviderDTO
	if err := c.Bind(&dto); err != nil {
		return res.BindFailed(err)
	}
	provider, err := h.service.Update(providerID, &dto)
	if err != nil {
		return res.Error(c, "UpdateProvider", err, "error updating provider")
	}
	return c.JSON(http.StatusOK, provider)
}
