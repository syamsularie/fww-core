package model

import "time"

type Flight struct {
	FlightID             int       `json:"flight_id"`
	FlightNumber         string    `json:"flight_number"`
	AirlineCode          string    `json:"airline_code"`
	DepartureAirportCode string    `json:"departure_airport_code"`
	ArrivalAirportCode   string    `json:"arrival_airport_code"`
	DepartureTime        time.Time `json:"departure_date_time"`
	ArrivalTime          time.Time `json:"arrival_date_time"`
	SeatCapacity         int       `json:"seat_capacity"`
}
