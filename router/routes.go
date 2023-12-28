package router

import (
	"fmt"
	"log"

	"firebase.google.com/go/v4/db"
	"github.com/gofiber/fiber/v2"
	"github.com/kristofkruller/calendar-service/handlers"
	"github.com/kristofkruller/calendar-service/models"
)

func SetupRoutes(app *fiber.App, db *db.Client) {
	events := app.Group("/events")

	events.Get("/:id", func(c *fiber.Ctx) error {
		eventId := c.Params(":id")
		event, err := handlers.GetOneEvent(db, eventId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(event)
	})

	events.Get("/", func(c *fiber.Ctx) error {
		events, err := handlers.GetAllEvents(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(events)
	})

	events.Post("/", func(c *fiber.Ctx) error {
		event := new(models.Event)
		if err := c.BodyParser(event); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		if err := handlers.SaveEvent(db, event); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(event)
	})

	events.Delete("/:id", func(c *fiber.Ctx) error {
		eventId := c.Params("id")
		if err := handlers.DeleteEvent(db, eventId); err != nil {
			log.Printf("error in deleting ref for user %s: %v\n", eventId, err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("Event's score with ID %s deleted successfully:)\n", eventId))
	})

	events.Delete("/", func(c *fiber.Ctx) error {
		if err := handlers.DeleteAllEvents(db); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return nil
	})
}
