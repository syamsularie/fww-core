package model

type Airline struct {
	AirlineID   int    `json:"airline_id"`
	AirlineCode string `json:"airline_code"`
	AirlineName string `json:"airline_name"`
}
