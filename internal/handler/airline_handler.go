package handler

import (
	"fww-core/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type Airline struct {
	AirlineUsecase usecase.AirlineExecutor
}

type AirlineHandler interface {
	GetAllAirlines(c *fiber.Ctx) error
}

func NewAirlineHandler(handler Airline) AirlineHandler {
	return &handler
}

func (handler *Airline) GetAllAirlines(c *fiber.Ctx) error {
	airlines, err := handler.AirlineUsecase.GetAllArlines()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(airlines)
}
