package model

type Flight struct {
	FlightNumber string `json:"flight_number"`
	AirlineCode  string `json:"airline_code"`
	Departure    string `json:"departure_airport_code"`
	Arrival      string `json:"arrival_airport_code"`
	SeatCapacity string `json:"seat_capacity"`
}
