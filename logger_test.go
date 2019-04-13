package logger

import (
	"fmt"
	"os"
)

func ExampleLoggerWithoutTimestamp() {
	// Most Verbose
	Level = 10
	// Timestamp Disabled
	TimeStamps = false
	// Color output enabled
	Color = true

	err := fmt.Errorf("It's a new Error")

	Debug("This is a lifeline for software engineers")
	Info("This helps software engineers while running application in prod")
	Warning("Something unusual happended, but application is still running")
	Error("Error %v", err)
	// Notice this does *NOT* exit!
	Critical("Something very bad happened, application should be stopped")
	Success("Hooray!! Application ran successfully")

	// Exit the function
	os.Exit(0)

	// Output:
	// 🐞  This is a lifeline for software engineers
	// 🧐  This helps software engineers while running application in prod
	// ⚠️  Something unusual happended, but application is still running
	// 😱  Error It's a new Error
	// 🚑  Something very bad happened, application should be stopped
	// ✅  Hooray!! Application ran successfully
}

func ExampleLoggerWithTimestamp() {
	// Most Verbose
	Level = 10
	// Timestamp Disabled
	TimeStamps = true
	// Color output enabled
	Color = true

	err := fmt.Errorf("It's a new Error")

	Debug("This is a lifeline for software engineers")
	Info("This helps software engineers while running application in prod")
	Warning("Something unusual happended, but application is still running")
	Error("Error %v", err)
	// Notice this does *NOT* exit!
	Critical("Something very bad happened, application should be stopped")
	Success("Hooray!! Application ran successfully")

	// Exit the function
	os.Exit(0)

	// Output:
	// 2019-04-13T13:39:38+05:30  🐞  This is a lifeline for software engineers
	// 2019-04-13T13:39:38+05:30  🧐  This helps software engineers while running application in prod
	// 2019-04-13T13:39:38+05:30  ⚠️  Something unusual happended, but application is still running
	// 2019-04-13T13:39:38+05:30  😱  Error It's a new Error
	// 2019-04-13T13:39:38+05:30  🚑  Something very bad happened, application should be stopped
	// 2019-04-13T13:39:38+05:30  ✅  Hooray!! Application ran successfully
}
