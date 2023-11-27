package model

type Seat struct {
	SeatID      int    `json:"seat_id"`
	SeatNumber  string `json:"seat_number"`
	FlightID    int    `json:"flight_i"`
	IsAvailable bool   `json:"is_available"`
}
