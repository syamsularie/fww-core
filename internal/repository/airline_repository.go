package repository

import (
	"database/sql"
	"fww-core/internal/model"
)

type AirlineRepository struct {
	DB *sql.DB
}

type AirlinePersister interface {
	GetAllAirlines() ([]model.Airline, error)
}

func NewAirlineRepository(airline AirlineRepository) AirlinePersister {
	return &airline
}

func (repo *AirlineRepository) GetAllAirlines() ([]model.Airline, error) {
	rows, err := repo.DB.Query("SELECT airline_id, airline_code, airline_name FROM airlines")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	airlines := make([]model.Airline, 0)
	for rows.Next() {
		var airline model.Airline
		if err := rows.Scan(&airline.AirlineID, &airline.AirlineCode, &airline.AirlineName); err != nil {
			return nil, err
		}
		airlines = append(airlines, airline)
	}

	return airlines, nil
}
