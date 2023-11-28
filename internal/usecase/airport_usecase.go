package usecase

import (
	"fww-core/internal/model"
	"fww-core/internal/repository"
)

type AirportUsecase struct {
	AirportRepo repository.AirportRepository
}

type AirportExecutor interface {
	GetAllAirports() ([]model.Airport, error)
}

func NewAirportUsecase(airportUsecase *AirportUsecase) AirportExecutor {
	return airportUsecase
}

func (uc *AirportUsecase) GetAllAirports() ([]model.Airport, error) {
	return uc.AirportRepo.GetAllAirports()
}
