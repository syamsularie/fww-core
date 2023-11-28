package repository

import (
	"database/sql"
	"fww-core/internal/model"
)

type AirportRepository struct {
	DB *sql.DB
}

type AirportPersister interface {
	GetAllAirports() ([]model.Airport, error)
}

func NewAirportRepository(airport AirportRepository) AirportPersister {
	return &airport
}

func (repo *AirportRepository) GetAllAirports() ([]model.Airport, error) {
	rows, err := repo.DB.Query("SELECT airport_id, airport_code, airport_name FROM airports")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	airports := make([]model.Airport, 0)
	for rows.Next() {
		var airport model.Airport
		if err := rows.Scan(&airport.AirportID, &airport.AirportCode, &airport.AirportName); err != nil {
			return nil, err
		}
		airports = append(airports, airport)
	}

	return airports, nil
}
