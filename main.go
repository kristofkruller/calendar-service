package main

import (
	"log"
)

func main() {
	log.Println("Starting the calendar service...")
	err := SetupAndRun()
	if err != nil {
		log.Fatalf("Failed to set up and run the application: %v", err)
	}
}
