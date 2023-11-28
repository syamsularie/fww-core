package handler

import (
	"fww-core/internal/model"
	"fww-core/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Passenger struct {
	PassengerUsecase usecase.PassengerExecutor
}

type PassengerHandler interface {
	GetPassengerById(c *fiber.Ctx) error
	CreatePassenger(c *fiber.Ctx) error
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

func (handler *Passenger) CreatePassenger(c *fiber.Ctx) error {
	var passenger model.Passenger
	if err := c.BodyParser(&passenger); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	id, err := handler.PassengerUsecase.CreatePassenger(&passenger)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	passenger.PassengerID = id

	return c.Status(fiber.StatusCreated).JSON(passenger)
}
