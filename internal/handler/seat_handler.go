package handler

import (
	"fww-core/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Seat struct {
	SeatUsecase usecase.SeatExecutor
}

type SeatHandler interface {
	GetAvailableSeatByFlightId(c *fiber.Ctx) error
	GetAllSeatByFlightId(c *fiber.Ctx) error
}

func NewSeatHandler(handler Seat) SeatHandler {
	return &handler
}

func (handler *Seat) GetAvailableSeatByFlightId(c *fiber.Ctx) error {
	flightIdString := c.Params("flight_id")
	flightId, _ := strconv.Atoi(flightIdString)
	availableSeats, err := handler.SeatUsecase.GetAvailableSeatByFlightId(flightId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(availableSeats)
}

func (handler *Seat) GetAllSeatByFlightId(c *fiber.Ctx) error {
	flightIdString := c.Params("flight_id")
	flightId, _ := strconv.Atoi(flightIdString)
	seats, err := handler.SeatUsecase.GetAllSeatByFlightId(flightId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(seats)
}
