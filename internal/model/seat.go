package model

type Seat struct {
	SeatID      int    `json:"seat_id"`
	SeatNumber  string `json:"seat_number"`
	FlightID    int    `json:"flight_id"`
	IsAvailable bool   `json:"is_available"`
}

type PassengerSeat struct {
	PassengerSeatID int `json:"passenger_seat_id"`
	SeatID          int `json:"seat_id"`
	PassengerID     int `json:"passenger_id"`
}
