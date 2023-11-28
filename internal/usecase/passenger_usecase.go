package usecase

import (
	"fww-core/internal/model"
	"fww-core/internal/repository"
)

type PassengerUsecase struct {
	PassengerRepo repository.PassengerRepository
}

type PassengerExecutor interface {
	GetPassengerById(passengerId int) (*model.Passenger, error)
	CreatePassenger(passenger *model.Passenger) (int, error)
}

func NewPassengerUsecase(passengerUsecase *PassengerUsecase) PassengerExecutor {
	return passengerUsecase
}

func (uc *PassengerUsecase) GetPassengerById(passengerId int) (*model.Passenger, error) {
	return uc.PassengerRepo.GetPassengerById(passengerId)
}

func (uc *PassengerUsecase) CreatePassenger(passenger *model.Passenger) (int, error) {
	return uc.PassengerRepo.CreatePassenger(passenger)
}
