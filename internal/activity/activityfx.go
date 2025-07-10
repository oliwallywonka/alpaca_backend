package activityfx

import (
	"github.com/labstack/echo/v4"
	"github.com/oliwallywonka/alpaca_backend/internal/activity/handlers"
	"github.com/oliwallywonka/alpaca_backend/internal/activity/repositories"
	"github.com/oliwallywonka/alpaca_backend/internal/activity/services"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"activityfx",
	fx.Provide(
		repositories.New,
		services.New,
		handlers.New,
	),
	fx.Invoke(func(c *echo.Echo, handler *handlers.Handler) {
		// REGISTER ROUTES
		c.GET("/activities", handler.GetPaginated)
		c.GET("/activities/:key", handler.GetActivityByUniqueKey)
		c.POST("/activities", handler.Save)
		c.PUT("/activities/:id", handler.Update)

		c.POST("/activities/:activityID/images", handler.SetImage)
		c.DELETE("/activities/:activityID/images/:imageID", handler.DeleteImage)
	}),
)
