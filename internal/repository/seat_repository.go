package repository

import (
	"database/sql"
	"fww-core/internal/model"
)

type SeatRepository struct {
	DB *sql.DB
}

type SeatPersister interface {
	GetAvailableSeatByFlightId(flightId int) ([]model.Seat, error)
	GetAllSeatByFlightId(flightId int) ([]model.Seat, error)
}

func NewSeatRepository(seat SeatRepository) SeatPersister {
	return &seat
}

func (repo *SeatRepository) GetAvailableSeatByFlightId(flightId int) ([]model.Seat, error) {
	rows, err := repo.DB.Query("SELECT seat_id, seat_number, flight_id, is_available FROM seats WHERE flight_id = ? AND is_available = 1", flightId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	availableSeats := make([]model.Seat, 0)
	for rows.Next() {
		var availableSeat model.Seat
		if err := rows.Scan(&availableSeat.SeatID, &availableSeat.SeatNumber, &availableSeat.FlightID, &availableSeat.IsAvailable); err != nil {
			return nil, err
		}
		availableSeats = append(availableSeats, availableSeat)
	}

	return availableSeats, nil
}

func (repo *SeatRepository) GetAllSeatByFlightId(flightId int) ([]model.Seat, error) {
	rows, err := repo.DB.Query("SELECT seat_id, seat_number, flight_id, is_available FROM seats WHERE flight_id = ?", flightId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	Seats := make([]model.Seat, 0)
	for rows.Next() {
		var Seat model.Seat
		if err := rows.Scan(&Seat.SeatID, &Seat.SeatNumber, &Seat.FlightID, &Seat.IsAvailable); err != nil {
			return nil, err
		}
		Seats = append(Seats, Seat)
	}

	return Seats, nil
}
