package repository

import (
	"database/sql"
	"fww-core/internal/model"
	"log"
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
	GetFlightsByCriteria(departureAirportCode, arrivalAirportCode string, date string) ([]model.Flight, error)
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
		if err := rows.Scan(&flight.FlightNumber, &flight.AirlineCode, &flight.DepartureAirportCode, &flight.ArrivalAirportCode, &flight.SeatCapacity); err != nil {
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
	err := row.Scan(&flight.FlightNumber, &flight.AirlineCode, &flight.DepartureAirportCode, &flight.ArrivalAirportCode, &flight.SeatCapacity)
	if err != nil {
		return nil, err
	}

	return &flight, nil
}

// CreateFlight adds a new flight to the database
func (repo *FlightRepository) CreateFlight(flight *model.Flight) error {
	_, err := repo.DB.Exec("INSERT INTO flights (flight_number, airline_code, departure_airport_code, arrival_airport_code, seat_capacity) VALUES (?, ?, ?, ?)",
		flight.FlightNumber, flight.AirlineCode, flight.DepartureAirportCode, flight.ArrivalAirportCode, flight.SeatCapacity)
	if err != nil {
		return err
	}

	return nil
}

// UpdateFlight updates a flight in the database
func (repo *FlightRepository) UpdateFlight(flight *model.Flight) error {
	_, err := repo.DB.Exec("UPDATE flights SET departure_airport_code=?, arrival_airport_code=?, seat_capacity=? WHERE flight_number=?",
		flight.DepartureAirportCode, flight.ArrivalAirportCode, flight.SeatCapacity, flight.FlightNumber)
	if err != nil {
		return err
	}

	return nil
}

// DeleteFlight removes a flight from the database by its ID
func (repo *FlightRepository) DeleteFlight(id string) error {
	_, err := repo.DB.Exec("DELETE FROM flights WHERE flight_number = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *FlightRepository) GetFlightsByCriteria(departureAirportCode, arrivalAirportCode string, date string) ([]model.Flight, error) {
	log.Println(date)
	rows, err := repo.DB.Query(`
		SELECT flight_id, flight_number, airline_code, 
		departure_airport_code, arrival_airport_code, 
		departure_date_time, arrival_date_time, seat_capacity
		FROM flights 
		WHERE departure_airport_code = ? 
		AND arrival_airport_code = ?
		AND DATE(departure_date_time) = ?
		ORDER BY departure_date_time ASC;
	`, departureAirportCode, arrivalAirportCode, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	flights := make([]model.Flight, 0)
	for rows.Next() {
		var flight model.Flight
		if err := rows.Scan(&flight.FlightID, &flight.FlightNumber, &flight.AirlineCode,
			&flight.DepartureAirportCode, &flight.ArrivalAirportCode, &flight.DepartureTime,
			&flight.ArrivalTime, &flight.SeatCapacity,
		); err != nil {
			return nil, err
		}
		flights = append(flights, flight)
	}

	return flights, nil
}
