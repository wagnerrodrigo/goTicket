package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wagnerrodrigo/goTicket/handlers"
	"github.com/wagnerrodrigo/goTicket/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "GoTicket",
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
