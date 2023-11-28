package usecase

import (
	"fww-core/internal/model"
	"fww-core/internal/repository"
)

// FlightUsecase represents the use case for managing flights
type FlightUsecase struct {
	FlightRepo repository.FlightRepository
}

type FlightExecutor interface {
	GetAllFlights() ([]model.Flight, error)
	GetFlightByID(id string) (*model.Flight, error)
	CreateFlight(flight *model.Flight) (string, error)
	UpdateFlight(flight *model.Flight) error
	DeleteFlight(id string) error
	GetFlightsByCriteria(departureAirportCode, arrivalAirportCode string, date string) ([]model.Flight, error)
}

// NewFlightUsecase creates a new instance of FlightUsecase
func NewFlightUsecase(flightUsecase *FlightUsecase) FlightExecutor {
	return flightUsecase
}

// GetAllFlights retrieves all flights
func (uc *FlightUsecase) GetAllFlights() ([]model.Flight, error) {
	return uc.FlightRepo.GetAllFlights()
}

// GetFlightByID retrieves a flight by its ID
func (uc *FlightUsecase) GetFlightByID(id string) (*model.Flight, error) {
	return uc.FlightRepo.GetFlightByID(id)
}

// CreateFlight adds a new flight
func (uc *FlightUsecase) CreateFlight(flight *model.Flight) (string, error) {

	err := uc.FlightRepo.CreateFlight(flight)
	if err != nil {
		return "", err
	}
	return "success", nil
}

// UpdateFlight updates a flight
func (uc *FlightUsecase) UpdateFlight(flight *model.Flight) error {

	return uc.FlightRepo.UpdateFlight(flight)
}

// DeleteFlight removes a flight by its ID
func (uc *FlightUsecase) DeleteFlight(id string) error {

	return uc.FlightRepo.DeleteFlight(id)
}

func (uc *FlightUsecase) GetFlightsByCriteria(departureAirportCode, arrivalAirportCode string, date string) ([]model.Flight, error) {
	return uc.FlightRepo.GetFlightsByCriteria(departureAirportCode, arrivalAirportCode, date)
}
