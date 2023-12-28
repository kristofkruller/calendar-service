package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kristofkruller/calendar-service/calendardb"
	"github.com/kristofkruller/calendar-service/config"
	"github.com/kristofkruller/calendar-service/router"
)

func SetupAndRun() error {
	// load env
	err := config.LoadENV()
	if err != nil {
		return err
	}

	// db
	calendardb.InitFirebase()

	// app
	app := fiber.New()

	// attach middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	// setup routes
	router.SetupRoutes(app, calendardb.DbClient)

	// get the port and start
	port := os.Getenv("PORT")
	app.Listen(":" + port)

	return nil
}
