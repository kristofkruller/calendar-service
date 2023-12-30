package main

import (
	"github.com/kristofkruller/calendar-service/app"
	"log"
)

func main() {
	log.Println("Starting the calendar service...")

	err := app.SetupAndRun()
	if err != nil {
		log.Fatalf("Failed to set up and run the application: %v", err)
	}

	log.Println("Calendar service is running successfully.")
}
