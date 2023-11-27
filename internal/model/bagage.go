package model

type Bagage struct {
	BagageID      int     `json:"bagage_id"`
	PassengerName string  `json:"passenger_name"`
	FlightID      int     `json:"flight_id"`
	WeightKG      float32 `json:"weight_kg"`
}
