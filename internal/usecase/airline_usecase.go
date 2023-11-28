package usecase

import (
	"fww-core/internal/model"
	"fww-core/internal/repository"
)

type AirlineUsecase struct {
	AirlineRepo repository.AirlineRepository
}

type AirlineExecutor interface {
	GetAllArlines() ([]model.Airline, error)
}

func NewAirlineUsecase(airlineUsecase *AirlineUsecase) AirlineExecutor {
	return airlineUsecase
}

func (uc *AirlineUsecase) GetAllArlines() ([]model.Airline, error) {
	return uc.AirlineRepo.GetAllAirlines()
}
