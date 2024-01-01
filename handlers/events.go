package handlers

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strconv"

	"firebase.google.com/go/v4/db"
	"github.com/kristofkruller/calendar-service/models"
)

func GetOneEvent(client *db.Client, eventId string) (*models.Event, error) {
	ref := client.NewRef("events/" + eventId)
	var event models.Event
	if err := ref.Get(context.Background(), &event); err != nil {
		log.Printf("Error gettng event: %v\n %s", err, eventId)
		return nil, err
	}
	return &event, nil
}

func GetAllEvents(client *db.Client) (map[string]models.Event, error) {
	ref := client.NewRef("events")
	var events map[string]models.Event
	if err := ref.Get(context.Background(), &events); err != nil {
		log.Printf("Error getting events: %v\n", err)
		return nil, err
	}
	return events, nil
}

func SaveEvent(db *db.Client, event *models.Event) error {
	ref := db.NewRef("events")

	// CHECK TIME FIELDS FOR UTC
	times := []string{event.CDate, event.Begin, event.End}
	for _, t := range times {
		parsedTime, err := ParseT(t)
		if err != nil {
			return fmt.Errorf("error parsing time: %v", err)
		}
		if isUTC, err := IsUTCTime(*parsedTime); err != nil || !isUTC {
			if err != nil {
				return fmt.Errorf("error checking UTC: %v", err)
			}
			return fmt.Errorf("UTC time needed for dates")
		}
	}

	// CHECK FOR ALREADY EXISTING ONES
	var exEvent models.Event
	existingEventRef := ref.Child(strconv.Itoa(int(event.Id)))
	if err := existingEventRef.Get(context.Background(), &exEvent); err != nil {
		return fmt.Errorf("error checking for existing event: %v", err)
	}
	if exEvent.Id != 0 {
		return fmt.Errorf("event with ID %d already exists", event.Id)
	}

	// RM WHITESPACE, SPEC, FORMATTING
	v := reflect.ValueOf(event).Elem()
	tOfStr := reflect.TypeOf("")
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Type() == tOfStr && field.CanSet() {
			cleaned := CleanString(field.String())
			field.SetString(cleaned)
		}
	}

	// RETURN SET
	if err := existingEventRef.Set(context.Background(), event); err != nil {
		return fmt.Errorf("error saving event: %v", err)
	}

	return nil
}

func DeleteEvent(client *db.Client, eventId string) error {
	ref := client.NewRef("events/" + eventId)

	if err := ref.Delete(context.Background()); err != nil {
		log.Printf("Error deleting event: %v\n", err)
		return err
	}
	return nil
}

func DeleteAllEvents(client *db.Client) error {
	ref := client.NewRef("events")
	if err := ref.Delete(context.Background()); err != nil {
		log.Printf("Error deleting all events: %v\n", err)
		return err
	}
	return nil
}
