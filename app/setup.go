package app

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kristofkruller/calendar-service/calendardb"
	"github.com/kristofkruller/calendar-service/config"
	"github.com/kristofkruller/calendar-service/router"
)

func SetupAndRun() error {
	// Load environment variables
	if err := config.LoadENV(); err != nil {
		log.Printf("Error loading environment variables: %v\n", err)
		return err
	}

	// Initialize the database
	if _, err := calendardb.InitFirebase(); err != nil {
		log.Printf("Error initializing Firebase: %v\n", err)
		return err
	}

	// Create a new Fiber application
	newApp := fiber.New()

	// Attach middleware
	newApp.Use(recover.New())
	newApp.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	// Setup routes
	router.SetupRoutes(newApp, calendardb.DbClient)

	// Get the port and start the application
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Port not specified in the environment, defaulting to 3000")
		port = "3000" // default port
	}

	if err := newApp.Listen(":" + port); err != nil {
		log.Printf("Error starting the server on port %s: %v\n", port, err)
		return err
	}

	log.Printf("Server started on port %s\n", port)
	return nil
}
