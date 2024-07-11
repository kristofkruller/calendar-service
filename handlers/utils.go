package handlers

import (
	"context"
	"fmt"
	"strings"
	"time"
	"unicode"

	"firebase.google.com/go/v4/db"
	"github.com/kristofkruller/calendar-service/models"
)

func IsEventConflict(db *db.Client, newEvent *models.Event) (bool, error) {
	ref := db.NewRef("events")

	var events map[string]models.Event
	if err := ref.Get(context.Background(), &events); err != nil {
		return false, fmt.Errorf("failed to retrieve events: %w", err)
	}

	var parsedTimes []time.Time
	times := []string{newEvent.CDate, newEvent.Begin, newEvent.End}
	for _, t := range times {
		parsedTime, err := ParseT(t)
		if err != nil {
			return false, fmt.Errorf("error parsing time: %v", err)
		}
		parsedTimes = append(parsedTimes, *parsedTime)
	}

	newEventBegin, newEventEnd := parsedTimes[1], parsedTimes[2]
	for _, event := range events {
		eventBegin, err := ParseT(event.Begin)
		if err != nil {
			return false, fmt.Errorf("error parsing existing event begin time: %v", err)
		}
		eventEnd, err := ParseT(event.End)
		if err != nil {
			return false, fmt.Errorf("error parsing existing event end time: %v", err)
		}

		// Check for exact time conflicts or overlaps
		if (newEventBegin.Equal(*eventBegin) || newEventEnd.Equal(*eventEnd)) ||
			(newEventBegin.Before(*eventEnd) && newEventEnd.After(*eventBegin)) {
			return true, nil
		}
	}
	return false, nil
}

// frontend must send time as UTC eg. "2023-04-12T15:04:05Z"
func ParseT(timestamp string) (*time.Time, error) {
	// parse string to time.Time
	t, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
func IsUTCTime(t time.Time) (bool, error) {
	return t.Location() == time.UTC, nil
}

func CleanString(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}
