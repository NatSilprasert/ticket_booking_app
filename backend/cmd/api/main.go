package main

import (
	"fmt"

	"github.com/NatSilprasert/ticket_booking_app/config"
	"github.com/NatSilprasert/ticket_booking_app/db"
	"github.com/NatSilprasert/ticket_booking_app/handlers"
	"github.com/NatSilprasert/ticket_booking_app/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName: "TicketBooking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)
	
	// Routing
	server := app.Group("/api")
	
	// Handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}


