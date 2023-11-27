package repository

import (
	"database/sql"
	"fww-core/internal/model"
)

// FlightRepository represents a repository for managing flights
type FlightRepository struct {
	DB *sql.DB
}

type FlightPersister interface {
	GetAllFlights() ([]model.Flight, error)
	GetFlightByID(id string) (*model.Flight, error)
	CreateFlight(flight *model.Flight) error
	UpdateFlight(flight *model.Flight) error
	DeleteFlight(id string) error
}

// NewFlightRepository creates a new instance of FlightRepository
func NewFlightRepository(flight FlightRepository) FlightPersister {
	return &flight
}

// GetAllFlights retrieves all flights from the database
func (repo *FlightRepository) GetAllFlights() ([]model.Flight, error) {
	rows, err := repo.DB.Query("SELECT flight_number, airline_code, departure_airport_code, arrival_airport_code, seat_capacity FROM flights")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	flights := make([]model.Flight, 0)
	for rows.Next() {
		var flight model.Flight
		if err := rows.Scan(&flight.FlightNumber, &flight.AirlineCode, &flight.Departure, &flight.Arrival, &flight.SeatCapacity); err != nil {
			return nil, err
		}
		flights = append(flights, flight)
	}

	return flights, nil
}

// GetFlightByID retrieves a flight by its ID from the database
func (repo *FlightRepository) GetFlightByID(id string) (*model.Flight, error) {
	row := repo.DB.QueryRow("SELECT flight_number, airline_code, departure_airport_code, arrival_airport_code, seat_capacity FROM flights WHERE flight_number = ?", id)

	var flight model.Flight
	err := row.Scan(&flight.FlightNumber, &flight.AirlineCode, &flight.Departure, &flight.Arrival, &flight.SeatCapacity)
	if err != nil {
		return nil, err
	}

	return &flight, nil
}

// CreateFlight adds a new flight to the database
func (repo *FlightRepository) CreateFlight(flight *model.Flight) error {
	_, err := repo.DB.Exec("INSERT INTO flights (flight_number, airline_code, departure_airport_code, arrival_airport_code, seat_capacity) VALUES (?, ?, ?, ?)",
		flight.FlightNumber, flight.AirlineCode, flight.Departure, flight.Arrival, flight.SeatCapacity)
	if err != nil {
		return err
	}

	return nil
}

// UpdateFlight updates a flight in the database
func (repo *FlightRepository) UpdateFlight(flight *model.Flight) error {
	_, err := repo.DB.Exec("UPDATE flights SET departure_airport_code=$1, arrival_airport_code=$2, seat_capacity=$3 WHERE flight_number=$4",
		flight.Departure, flight.Arrival, flight.SeatCapacity, flight.FlightNumber)
	if err != nil {
		return err
	}

	return nil
}

// DeleteFlight removes a flight from the database by its ID
func (repo *FlightRepository) DeleteFlight(id string) error {
	_, err := repo.DB.Exec("DELETE FROM flights WHERE flight_number = $1", id)
	if err != nil {
		return err
	}

	return nil
}
