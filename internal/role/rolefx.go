package rolefx

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Module = fx.Module("rolefx",
	fx.Provide(
		NewHandler,
		NewService,
	),
	fx.Invoke(func(e *echo.Echo, h *Handler) {
		e.GET("/roles", h.GetRoles)
		e.GET("/roles/:roleID", h.GetRole)
		e.POST("/roles", h.Create)
		e.PUT("/roles/:roleID", h.Update)
	}),
)
