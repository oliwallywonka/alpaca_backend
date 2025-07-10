package resourcefx

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"resourcefx",
	fx.Provide(
		NewResourceProviderService,
		NewResourceService,
		NewHandler,
	),
	fx.Invoke(func(e *echo.Echo, h *ResourceHandler) {
		e.GET("/resources", h.GetResources)
		e.GET("/resourcesv2", h.GetResourcesV2)
		e.GET("/resources/:resourceID", h.GetResource)
		e.POST("/resources", h.Save)
		e.PUT("/resources/:resourceID", h.Update)
	}),
)
