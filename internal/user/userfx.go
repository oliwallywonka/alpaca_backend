package userfx

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	resourcefx "github.com/oliwallywonka/alpaca_backend/internal/resource"
)

var Module = fx.Module(
	"userfx",
	fx.Provide(
		NewHandler,
		NewService,
	),
	fx.Invoke(func(e *echo.Echo, h *Handler, hr *resourcefx.ResourceHandler) {
		e.GET("/users", h.GetUsers)
		e.GET("/users/:userID", h.GetUser)
		e.POST("/users", h.Create)
		e.PUT("/users/:userID", h.Update)

		e.GET("/users/:providerID/resources", hr.GetResourceProviders)
		e.GET("/users/:providerID/resources/:providerResourceID", hr.GetResourceProvider)
		e.POST("/users/:providerID/resources", hr.SaveResourceProvider)
		e.PUT("/users/:providerID/resources/:providerResourceID", hr.UpdateResourceProvider)
		e.DELETE("/users/:providerID/resources/:providerResourceID", hr.DeleteResourceProvider)
	}),
)
