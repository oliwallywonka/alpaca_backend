package mealfx

import (
	"github.com/labstack/echo/v4"
	"github.com/oliwallywonka/alpaca_backend/internal/meal/handlers"
	"github.com/oliwallywonka/alpaca_backend/internal/meal/repositories"
	"github.com/oliwallywonka/alpaca_backend/internal/meal/services"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"mealfx",
	fx.Provide(
		repositories.New,
		services.New,
		handlers.New,
	),
	fx.Invoke(func(e *echo.Echo, h *handlers.Handler) {
		// REGISTER ROUTES
		e.GET("/meals", h.GetPaginated)
		e.GET("/meals/:mealID", h.GetByUniqueKey)
		e.POST("/meals", h.SaveMeal)
		e.PUT("/meals/:mealID", h.UpdateMeal)
	}),
)
