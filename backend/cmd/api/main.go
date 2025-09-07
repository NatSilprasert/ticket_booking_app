package main

import (
	"github.com/NatSilprasert/ticket_booking_app/handlers"
	"github.com/NatSilprasert/ticket_booking_app/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "TicketBooking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(nil)
	
	// Routing
	server := app.Group("/api")
	
	// Handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}


