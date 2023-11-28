package usecase

import (
	"fww-core/internal/model"
	"fww-core/internal/repository"
)

type SeatUsecase struct {
	SeatRepo repository.SeatRepository
}

type SeatExecutor interface {
	GetAvailableSeatByFlightId(flightId int) ([]model.Seat, error)
	GetAllSeatByFlightId(flightId int) ([]model.Seat, error)
}

func NewSeatUsecase(seatUsecase *SeatUsecase) SeatExecutor {
	return seatUsecase
}

func (uc *SeatUsecase) GetAvailableSeatByFlightId(flightId int) ([]model.Seat, error) {
	return uc.SeatRepo.GetAvailableSeatByFlightId(flightId)
}

func (uc *SeatUsecase) GetAllSeatByFlightId(flightId int) ([]model.Seat, error) {
	return uc.SeatRepo.GetAllSeatByFlightId(flightId)
}
