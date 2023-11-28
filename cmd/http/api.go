package main

import (
	"fmt"
	"fww-core/config"
	"fww-core/config/middleware"
	"fww-core/internal/handler"
	"fww-core/internal/repository"
	"fww-core/internal/usecase"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	baseDep := config.NewBaseDep()
	loadEnv(baseDep.Logger)
	db, err := config.NewDbPool(baseDep.Logger)

	if err != nil {
		os.Exit(1)
	}

	dbCollector := middleware.NewStatsCollector("fww", db)
	prometheus.MustRegister(dbCollector)
	fiberProm := middleware.NewWithRegistry(prometheus.DefaultRegisterer, "fww-core", "", "", map[string]string{})

	//=== repository lists start ===//
	flightRepo := repository.FlightRepository(repository.FlightRepository{
		DB: db,
	})

	airlineRepo := repository.AirlineRepository(repository.AirlineRepository{
		DB: db,
	})

	airportRepo := repository.AirportRepository(repository.AirportRepository{
		DB: db,
	})

	seatRepo := repository.SeatRepository(repository.SeatRepository{
		DB: db,
	})
	//=== repository lists end ===//

	//=== usecase lists start ===//
	flightUsecase := usecase.NewFlightUsecase(&usecase.FlightUsecase{
		FlightRepo: flightRepo,
	})

	airlineUsecase := usecase.NewAirlineUsecase(&usecase.AirlineUsecase{
		AirlineRepo: airlineRepo,
	})

	airportUsecase := usecase.NewAirportUsecase(&usecase.AirportUsecase{
		AirportRepo: airportRepo,
	})

	seatUsecase := usecase.NewSeatUsecase(&usecase.SeatUsecase{
		SeatRepo: seatRepo,
	})
	//=== usecase lists end ===//

	//=== handler lists start ===//
	flightHandler := handler.NewFlightHandler(handler.Flight{
		FlightUsecase: flightUsecase,
	})

	airlineHandler := handler.NewAirlineHandler(handler.Airline{
		AirlineUsecase: airlineUsecase,
	})

	airportHandler := handler.NewAirportHandler(handler.Airport{
		AirportUsecase: airportUsecase,
	})

	seatHandler := handler.NewSeatHandler(handler.Seat{
		SeatUsecase: seatUsecase,
	})
	//=== handler lists end ===//
	app := fiber.New(fiber.Config{
		BodyLimit: 30 * 1024 * 1024,
	})

	app.Use(fiberProm.Middleware)
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(pprof.New())
	app.Use(logger.New(logger.Config{
		Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeInterval: time.Millisecond,
		TimeFormat:   "02-01-2006 15:04:05",
		TimeZone:     "Indonesia/Jakarta",
	}))
	// Define a route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	//=== healthz route
	// app.Get("/", Healthz)
	app.Get("/healthz", Healthz)

	//Flight Routes
	app.Get("/flights", flightHandler.GetAllFlights)
	app.Get("/flights/:id", flightHandler.GetFlightByID)
	app.Post("/flights", flightHandler.CreateFlight)
	app.Put("/flights/:id", flightHandler.UpdateFlight)
	app.Delete("/flights/:id", flightHandler.DeleteFlight)

	//Airline Routes
	app.Get("/airlines", airlineHandler.GetAllAirlines)

	//Airport Routes
	app.Get("/airports", airportHandler.GetAllAirports)

	//Seat Routes
	app.Get("/seats/available/:flight_id", seatHandler.GetAvailableSeatByFlightId)
	app.Get("/seats/:flight_id", seatHandler.GetAllSeatByFlightId)

	//=== listen port ===//
	if err := app.Listen(fmt.Sprintf(":%s", "3000")); err != nil {
		log.Fatal(err)
	}

}

func Healthz(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Service is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}

func loadEnv(logger config.Logger) {
	_, err := os.Stat(".env")
	if err == nil {
		err = godotenv.Load()
		if err != nil {
			logger.Error("no .env files provided")
		}
	}
}
