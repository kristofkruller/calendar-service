package main

import "github.com/kristofkruller/calendar-service/app"

func main() {
	err := app.SetupAndRun()
	if err != nil {
		panic(err)
	}
}
