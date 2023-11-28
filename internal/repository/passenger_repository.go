package repository

import (
	"database/sql"
	"fww-core/internal/model"
)

type PassengerRepository struct {
	DB *sql.DB
}

type PassengerPersister interface {
	GetPassengerById(passengerId int) (*model.Passenger, error)
}

func (repo *PassengerRepository) GetPassengerById(passengerId int) (*model.Passenger, error) {
	row := repo.DB.QueryRow("SELECT passenger_id, username, email, first_name, last_name, phone_number FROM passengers WHERE passenger_id = ?", passengerId)

	var passenger model.Passenger
	err := row.Scan(&passenger.PassengerID, &passenger.Username, &passenger.Email, &passenger.FirstName, &passenger.LastName, &passenger.Phone)
	if err != nil {
		return nil, err
	}

	return &passenger, nil
}
