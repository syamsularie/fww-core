package handler

import (
	"fww-core/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Passenger struct {
	PassengerUsecase usecase.PassengerExecutor
}

type PassengerHandler interface {
	GetPassengerById(c *fiber.Ctx) error
}

func NewPassengerHandler(handler Passenger) PassengerHandler {
	return &handler
}

func (handler *Passenger) GetPassengerById(c *fiber.Ctx) error {
	passengerIdString := c.Params("id")
	passengerId, _ := strconv.Atoi(passengerIdString)

	passenger, err := handler.PassengerUsecase.GetPassengerById(passengerId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(passenger)
}
