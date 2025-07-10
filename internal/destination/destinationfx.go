package destinationfx

import (
	"go.uber.org/fx"

	"github.com/labstack/echo/v4"
)

var Module = fx.Module("destination",
	fx.Provide(
		NewHandler,
		NewService,
	),
	fx.Invoke(func(e *echo.Echo, h *Handler) {
		//group := e.Group("/destinations")
		e.GET("/destinations", h.GetAll)
		e.GET("/destinations/:destinationID", h.GetByID)
		e.POST("/destinations", h.Create)
		e.PUT("/destinations/:destinationID", h.Update)
	}),
)
