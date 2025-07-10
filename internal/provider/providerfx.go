package providerfx

import (
	"github.com/labstack/echo/v4"
	//servicefx "github.com/oliwallywonka/alpaca_backend/internal/service"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"providerfx",
	fx.Provide(
		NewHandler,
		NewService,
	),
	fx.Invoke(func(e *echo.Echo, h *Handler, /* hs *servicefx.ServiceHandler */) {
		e.GET("/providers", h.GetAll)
		e.GET("/providers/:providerID", h.GetProvider)
		e.POST("/providers", h.Create)
		e.PUT("/providers/:providerID", h.Update)

		/* e.GET("/providers/:providerID/services", hs.GetProviderServices)
		e.POST("/providers/:providerID/services", hs.SaveProviderService)
		e.PUT("/providers/:providerID/services/:providerServiceID", hs.UpdateProviderService)
		e.DELETE("/providers/:providerID/services/:providerServiceID", hs.DeleteProviderService) */
	}),
)
