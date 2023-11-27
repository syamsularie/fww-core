package handler

import (
	"fww-core/internal/model"
	"fww-core/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type Flight struct {
	FlightUsecase usecase.FlightExecutor
}

type FlightaHandler interface {
	GetAllFlights(c *fiber.Ctx) error
	GetFlightByID(c *fiber.Ctx) error
	CreateFlight(c *fiber.Ctx) error
	UpdateFlight(c *fiber.Ctx) error
	DeleteFlight(c *fiber.Ctx) error
}

func NewFlightHandler(handler Flight) FlightaHandler {
	return &handler
}

func (handler *Flight) GetAllFlights(c *fiber.Ctx) error {
	flights, err := handler.FlightUsecase.GetAllFlights()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(flights)
}

func (handler *Flight) GetFlightByID(c *fiber.Ctx) error {
	id := c.Params("id")

	flight, err := handler.FlightUsecase.GetFlightByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if flight == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Flight not found"})
	}

	return c.JSON(flight)
}

func (handler *Flight) CreateFlight(c *fiber.Ctx) error {
	var flight model.Flight
	if err := c.BodyParser(&flight); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	id, err := handler.FlightUsecase.CreateFlight(&flight)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	flight.FlightNumber = id

	return c.Status(fiber.StatusCreated).JSON(flight)
}

func (handler *Flight) UpdateFlight(c *fiber.Ctx) error {
	id := c.Params("id")

	var flight model.Flight
	if err := c.BodyParser(&flight); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	flight.FlightNumber = id
	err := handler.FlightUsecase.UpdateFlight(&flight)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(flight)
}

func (handler *Flight) DeleteFlight(c *fiber.Ctx) error {
	id := c.Params("id")

	err := handler.FlightUsecase.DeleteFlight(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
}
