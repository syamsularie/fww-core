CREATE TABLE passengers (
    passenger_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    ktp VARCHAR(20) NOT NULL,
    phone_number VARCHAR(15),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE airports (
    airport_id INT AUTO_INCREMENT PRIMARY KEY,
    airport_code VARCHAR(3),
    airport_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE airlines (
    airline_id INT AUTO_INCREMENT PRIMARY KEY,
    airline_code VARCHAR(3),
    airline_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE flights (
    flight_id INT AUTO_INCREMENT PRIMARY KEY,
    flight_number VARCHAR(10),
    airline_code VARCHAR(3),
    departure_airport_code VARCHAR(3),
    arrival_airport_code VARCHAR(3),
    departure_date_time DATETIME,
    arrival_date_time DATETIME,
    seat_capacity INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE seats (
    seat_id INT AUTO_INCREMENT PRIMARY KEY,
    seat_number VARCHAR(3),
    flight_id INT,
    is_available BOOLEAN,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE passenger_seats (
    passenger_seat_id INT AUTO_INCREMENT PRIMARY KEY,
    seat_id INT ,
    passenger_id INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS baggage (
    bagage_id INT AUTO_INCREMENT PRIMARY KEY,
    passenger_name VARCHAR(255),
    flight_id INT,
    weight_kg DECIMAL(5, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- data airports
INSERT INTO airports (airport_code, airport_name) VALUES ("UPG", "Makassar");
INSERT INTO airports (airport_code, airport_name) VALUES ("CGK", "Jakarta");
INSERT INTO airports (airport_code, airport_name) VALUES ("JYP", "Jayapura");
INSERT INTO airports (airport_code, airport_name) VALUES ("SMG", "Semarang");

-- data airlines
INSERT INTO airlines (airline_code, airline_name) VALUES ("SG1", "Singa Air");
INSERT INTO airlines (airline_code, airline_name) VALUES ("AY1", "Ayam Air");
INSERT INTO airlines (airline_code, airline_name) VALUES ("BB1", "Bebek Air");


-- data flights
INSERT INTO flights (flight_number, airline_code, departure_airport_code, arrival_airport_code, departure_date_time, arrival_date_time, seat_capacity) VALUES ("JT2", "SG1", "SMG", "CGK", "2023-11-23 21:50:39", "2023-11-23 23:50:39", 200);
INSERT INTO flights (flight_number, airline_code, departure_airport_code, arrival_airport_code, departure_date_time, arrival_date_time, seat_capacity) VALUES ("UD1", "AY1", "UPG", "CGK", "2023-11-23 21:50:39", "2023-11-23 23:50:39", 200);
INSERT INTO flights (flight_number, airline_code, departure_airport_code, arrival_airport_code, departure_date_time, arrival_date_time, seat_capacity) VALUES ("JT2", "SG1", "JYP", "UPG", "2023-11-23 21:50:39", "2023-11-23 23:50:39", 200);


-- ALTER TABLE flights ADD CONSTRAINT fk_airline_code FOREIGN KEY (airline_code) REFERENCES airlines(airline_code);
-- ALTER TABLE flights ADD CONSTRAINT fk_departure_airport_code FOREIGN KEY (departure_airport_code) REFERENCES airports(airport_code);
-- ALTER TABLE flights ADD CONSTRAINT fk_arival_airport_code FOREIGN KEY (arrival_airport_code) REFERENCES airports(airport_code);

-- ALTER TABLE seats ADD CONSTRAINT fk_flight_number2 FOREIGN KEY (flight_number) REFERENCES flights(flight_number);

-- ALTER TABLE passenger_seats ADD CONSTRAINT fk_passenger_seat FOREIGN KEY (seat_id) REFERENCES seat(seat_id);
