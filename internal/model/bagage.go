package model

type Bagage struct {
	BagageID    int     `json:"bagage_id"`
	PassengerID int     `json:"passenger_id"`
	FlightID    int     `json:"flight_id"`
	WeightKG    float32 `json:"weight_kg"`
}
