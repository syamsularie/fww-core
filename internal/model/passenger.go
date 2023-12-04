package model

type Passenger struct {
	PassengerID int    `json:"passenger_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Ktp         string `json:"ktp"`
	Phone       string `json:"phone_number"`
	Username    string `json:"username"`
}
