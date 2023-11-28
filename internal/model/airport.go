package model

type Airport struct {
	AirportID   int    `json:"airport_id"`
	AirportCode string `json:"airport_code"`
	AirportName string `json:"airport_name"`
}
