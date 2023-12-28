package handlers

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/db"
	"github.com/kristofkruller/calendar-service/models"
)

func GetOneEvent(client *db.Client, eventId string) (*models.Event, error) {
	ref := client.NewRef("events/" + eventId)
	var event models.Event
	if err := ref.Get(context.Background(), &event); err != nil {
		return nil, err
	}
	return &event, nil
}

func GetAllEvents(client *db.Client) (map[string]models.Event, error) {
	ref := client.NewRef("events")
	var events map[string]models.Event
	if err := ref.Get(context.Background(), &events); err != nil {
		return nil, err
	}
	return events, nil
}

func SaveEvent(client *db.Client, event *models.Event) error {
	ref := client.NewRef("events/" + fmt.Sprint(event.Id))
	if err := ref.Set(context.Background(), event); err != nil {
		return err
	}
	return nil
}

func DeleteEvent(client *db.Client, eventId string) error {
	ref := client.NewRef("events/" + eventId)

	if err := ref.Delete(context.Background()); err != nil {
		return err
	}
	return nil
}

func DeleteAllEvents(client *db.Client) error {
	ref := client.NewRef("events")
	if err := ref.Delete(context.Background()); err != nil {
		return err
	}
	return nil
}
