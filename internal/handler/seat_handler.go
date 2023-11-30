package handler

import (
	"fww-core/internal/model"
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
	ReserveSeat(c *fiber.Ctx) error
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

func (handler *Seat) ReserveSeat(c *fiber.Ctx) error {
	var passengerSeat model.PassengerSeat
	if err := c.BodyParser(&passengerSeat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	seatIdString := c.Params("id")
	seatId, _ := strconv.Atoi(seatIdString)

	passengerSeat.SeatID = seatId

	id, err := handler.SeatUsecase.SavePassengerSeats(passengerSeat.SeatID, passengerSeat.PassengerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	passengerSeat.PassengerSeatID = id

	return c.Status(fiber.StatusCreated).JSON(passengerSeat)
}
