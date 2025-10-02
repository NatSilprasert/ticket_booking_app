package main

import (
	"fmt"

	"github.com/NatSilprasert/ticket_booking_app/config"
	"github.com/NatSilprasert/ticket_booking_app/db"
	"github.com/NatSilprasert/ticket_booking_app/handlers"
	"github.com/NatSilprasert/ticket_booking_app/middlewares"
	"github.com/NatSilprasert/ticket_booking_app/repositories"
	"github.com/NatSilprasert/ticket_booking_app/services"
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
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)
	
	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(db))

	// Handlers
	handlers.NewEventHandler(privateRoutes.Group("/event"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)

	app.Listen(fmt.Sprintf(":%s", envConfig.ServerPort))
}


