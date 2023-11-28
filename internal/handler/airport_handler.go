package handler

import (
	"fww-core/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type Airport struct {
	AirportUsecase usecase.AirportExecutor
}

type AirportHandler interface {
	GetAllAirports(c *fiber.Ctx) error
}

func NewAirportHandler(handler Airport) AirportHandler {
	return &handler
}

func (handler *Airport) GetAllAirports(c *fiber.Ctx) error {
	airports, err := handler.AirportUsecase.GetAllAirports()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(airports)
}
