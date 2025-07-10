package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/oliwallywonka/alpaca_backend/internal/hotel/dtos"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/internal/shared/utils"
)

func (h *Handler) GetHotelRooms(c echo.Context) error {
	res := shared.NewAPI()
	hotelID := c.Param("hotelID")

	hotelRooms, err := h.s.GetHotelRooms(hotelID)
	if err != nil {
		return res.Error(c, "GetHotelRooms", err, "error getting hotel rooms")
	}
	return c.JSON(res.Ok(hotelRooms))
}

func (h *Handler) CreateHotelRoom(c echo.Context) error {
	res := shared.NewAPI()
	dto := dtos.CreateHotelRoomDTO{}

	if err := c.Bind(dto); err != nil {
		return res.BindFailed(err)
	}

	if errors := utils.ValidateRequest(dto); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.s.CreateHotelRoom(&dto, c.Param("hotelID"))
	if err != nil {
		return res.Error(c, "CreateHotelRoom", err, "error creating hotel room")
	}
	return c.JSON(res.Created(nil))
}

func (h *Handler) UpdateHotelRoom(c echo.Context) error {
	res := shared.NewAPI()
	hotelRoomID := c.Param("hotelRoomID")
	dto := dtos.UpdateHotelRoomDTO{}

	if err := c.Bind(dto); err != nil {
		return res.BindFailed(err)
	}

	if errors := utils.ValidateRequest(dto); len(errors) > 0 {
		return res.Error(c, "ValidateRequest", nil, errors...)
	}

	err := h.s.UpdateHotelRoom(&dto, hotelRoomID)
	if err != nil {
		return res.Error(c, "UpdateHotelRoom", err, "error updating hotel room")
	}
	return c.JSON(res.Updated(nil))
}

func (h *Handler) DeleteHotelRoom(c echo.Context) error {
	res := shared.NewAPI()

	hotelRoomID := c.Param("hotelRoomID")
	err := h.s.DeleteHotelRoom(hotelRoomID)
	if err != nil {
		return res.Error(c, "DeleteHotelRoom", err, "error deleting hotel room")
	}
	return c.JSON(res.Deleted(nil))
}
