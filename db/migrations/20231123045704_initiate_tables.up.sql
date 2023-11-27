CREATE TABLE passengers (
    passenger_id BIGINT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(15),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE airports (
    airport_code VARCHAR(3) PRIMARY KEY,
    airport_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE airlines (
    airline_code VARCHAR(3) PRIMARY KEY,
    airline_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE flights (
    flight_number VARCHAR(10) PRIMARY KEY,
    airline_code VARCHAR(3),
    departure_airport_code VARCHAR(3),
    arrival_airport_code VARCHAR(3),
    departure_date_time DATETIME,
    arrival_date_time DATETIME,
    seat_capacity INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE reservations (
    reservation_id INT PRIMARY KEY,
    flight_number VARCHAR(10),
    passenger_id BIGINT,
    reservation_time DATETIME,
    seat_number INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
    
CREATE TABLE seats (
    seat_number INT PRIMARY KEY,
    flight_number VARCHAR(10),
    is_available BOOLEAN,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE payments (
    payment_id SERIAL PRIMARY KEY,
    reservation_id INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    payment_status VARCHAR(20) NOT NULL,
    payment_method VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS baggage (
    id INT AUTO_INCREMENT PRIMARY KEY,
    passenger_name VARCHAR(255),
    flight_id INT,
    weight_kg DECIMAL(5, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE flights ADD CONSTRAINT fk_airline_code FOREIGN KEY (airline_code) REFERENCES airlines(airline_code);
ALTER TABLE flights ADD CONSTRAINT fk_departure_airport_code FOREIGN KEY (departure_airport_code) REFERENCES airports(airport_code);
ALTER TABLE flights ADD CONSTRAINT fk_arival_airport_code FOREIGN KEY (arrival_airport_code) REFERENCES airports(airport_code);

ALTER TABLE reservations ADD CONSTRAINT fk_flight_number FOREIGN KEY (flight_number) REFERENCES flights(flight_number);
ALTER TABLE reservations ADD CONSTRAINT fk_passenger_id FOREIGN KEY (passenger_id) REFERENCES passengers(passenger_id);

ALTER TABLE payments ADD CONSTRAINT fk_reservation_id FOREIGN KEY (reservation_id) REFERENCES reservations(reservation_id);

ALTER TABLE seats ADD CONSTRAINT fk_flight_number2 FOREIGN KEY (flight_number) REFERENCES flights(flight_number);

