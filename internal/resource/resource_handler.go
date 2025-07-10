package resourcefx

import (
	"github.com/labstack/echo/v4"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	"github.com/oliwallywonka/alpaca_backend/internal/core"
	"github.com/oliwallywonka/alpaca_backend/internal/core/search"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/shared/utils"
)

type ResourceHandler struct {
	r  ResourceService
	rp ResourceProviderService
}

func NewHandler(r ResourceService, rp ResourceProviderService) *ResourceHandler {
	return &ResourceHandler{
		r:  r,
		rp: rp,
	}
}

func (h *ResourceHandler) GetResources(c echo.Context) error {
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
		return res.BindFailed(err)
	}
	if errors := utils.ValidateRequest(dto); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}
	response, err := h.r.GetPaginated(dto)
	if err != nil {
		return res.Error(c, "GetPaginated", err, "error getting paginated resources")
	}
	return c.JSON(res.Ok(response))
}

func (h *ResourceHandler) GetResourcesV2(c echo.Context) error {
	db, err := core.NewDB()
	if err != nil {
		return c.JSON(400, err.Error())
	}
	fieldsResolver := search.NewSimpleFieldResolver(
		"id", "name", "description", "location", "created_at", "updated_at",
		"resource_resource_type.id",
	)
	baseQuery := db.
		Select("resource.*").
		From("resource")
	provider := search.NewProvider(fieldsResolver).Query(baseQuery)
	resources := []*struct {
		model.Resource
		ResourceResourceType model.ResourceResourceType
	}{}
	response, err := provider.ParseAndExec(c.QueryString(), &resources)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, response)
}

func (h *ResourceHandler) GetResource(c echo.Context) error {
	res := shared.NewAPI()
	resourceID := c.Param("resourceID")
	resource, err := h.r.GetByUniqueKey(resourceID)
	if err != nil {
		return res.Error(c, "GetByUniqueKey", err, "error getting resource")
	}
	return c.JSON(res.Ok(resource))
}

func (h *ResourceHandler) Save(c echo.Context) error {
	res := shared.NewAPI()
	resource := &model.Resource{}
	if err := c.Bind(resource); err != nil {
		return res.BindFailed(err)
	}
	response, err := h.r.Save(resource)
	if err != nil {
		return res.Error(c, "Save", err, "error saving resource")
	}
	return c.JSON(res.Ok(response))
}

func (h *ResourceHandler) Update(c echo.Context) error {
	res := shared.NewAPI()
	resourceID := c.Param("resourceID")
	var dto UpdateResourceDTO
	if err := c.Bind(&dto); err != nil {
		return res.BindFailed(err)
	}
	response, err := h.r.Update(resourceID, &dto)
	if err != nil {
		return res.Error(c, "Update", err, "error updating resource")
	}
	return c.JSON(res.Ok(response))
}

func (h *ResourceHandler) GetResourceProviders(c echo.Context) error {
	res := shared.NewAPI()
	providerID := c.Param("providerID")
	providerResources, err := h.rp.GetByProviderID(providerID)
	if err != nil {
		return res.Error(c, "GetResourceProviders", err, "error getting provider resources")
	}
	return c.JSON(res.Ok(providerResources))
}

func (h *ResourceHandler) GetResourceProvider(c echo.Context) error {
	res := shared.NewAPI()
	providerResourceID := c.Param("providerResourceID")
	providerResource, err := h.rp.GetByID(providerResourceID)
	if err != nil {
		return res.Error(c, "GetResourceProvider", err, "error getting provider resource")
	}
	return c.JSON(res.Ok(providerResource))
}

func (h *ResourceHandler) SaveResourceProvider(c echo.Context) error {
	res := shared.NewAPI()
	var model model.ResourceProvider
	if err := c.Bind(&model); err != nil {
		return res.BindFailed(err)
	}
	response, err := h.rp.Save(&model)
	if err != nil {
		return res.Error(c, "SaveResourceProvider", err, "error saving provider resource")
	}
	return c.JSON(res.Ok(response))
}

func (h *ResourceHandler) UpdateResourceProvider(c echo.Context) error {
	res := shared.NewAPI()
	providerResourceID := c.Param("providerResourceID")
	var model model.ResourceProvider
	if err := c.Bind(&model); err != nil {
		return res.BindFailed(err)
	}
	response, err := h.rp.Update(providerResourceID, &model)
	if err != nil {
		return res.Error(c, "UpdateResourceProviders", err, "error updating provider resource")
	}
	return c.JSON(res.Ok(response))
}

func (h *ResourceHandler) DeleteResourceProvider(c echo.Context) error {
	res := shared.NewAPI()
	providerResourceID := c.Param("providerResourceID")
	err := h.rp.Delete(providerResourceID)
	if err != nil {
		return res.Error(c, "DeleteResourceProvider", err, "error deleting provider resource")
	}
	return c.JSON(res.Ok(nil))
}
