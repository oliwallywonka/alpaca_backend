package hotelfx

import (
	"go.uber.org/fx"
	"github.com/labstack/echo/v4"

	"github.com/oliwallywonka/alpaca_backend/internal/hotel/handlers"
	"github.com/oliwallywonka/alpaca_backend/internal/hotel/repositories"
	"github.com/oliwallywonka/alpaca_backend/internal/hotel/services"
)

var Module = fx.Module(
	"hotelfx",
	fx.Provide(
		repositories.New,
		services.New,
		handlers.New,
	),
	fx.Invoke(func(e *echo.Echo, h *handlers.Handler) {
		// REGISTER ROUTES
		e.GET("/hotels", h.GetPaginated)
		e.GET("/hotels/:key", h.GetByUniqueKey)
		e.POST("/hotels", h.SaveHotel)
		e.PUT("/hotels/:id", h.UpdateHotel)

		e.GET("/hotels/:hotelID/rooms", h.GetHotelRooms)
		e.POST("/hotels/:hotelID/rooms", h.CreateHotelRoom)
		e.PUT("/hotels/:hotelID/rooms/:hotelRoomID", h.UpdateHotelRoom)
		e.DELETE("/hotels/:hotelID/rooms/:hotelRoomID", h.DeleteHotelRoom)
	}),
)
