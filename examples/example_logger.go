package logger_test

import (
	"fmt"
	"os"

	logger "github.com/vaibhavsingh97/emoji-logger"
)

// ExampleLoggerWithoutTimestamp function demos logger package functionality.
func ExampleLoggerWithoutTimestamp() {
	// Most Verbose
	logger.Level = 10
	// Timestamp Disabled
	logger.TimeStamps = false
	// Color output enabled
	logger.Color = true

	err := fmt.Errorf("It's a new Error")

	logger.Debug("This is a lifeline for software engineers")
	logger.Info("This helps software engineers while running application in prod")
	logger.Warning("Something unusual happended, but application is still running")
	logger.Error("Error %v", err)
	// Note: this does NOT exit!
	logger.Critical("Something very bad happened, application should be stopped")
	logger.Success("Hooray!! Application ran successfully")

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

// ExampleLoggerWithTimestamp function demos logger package functionality.
func ExampleLoggerWithTimestamp() {
	// Most Verbose
	logger.Level = 10
	// Timestamp Disabled
	logger.TimeStamps = true
	// Color output enabled
	logger.Color = true

	err := fmt.Errorf("It's a new Error")

	logger.Debug("This is a lifeline for software engineers")
	logger.Info("This helps software engineers while running application in prod")
	logger.Warning("Something unusual happended, but application is still running")
	logger.Error("Error %v", err)
	// Note: this does NOT exit!
	logger.Critical("Something very bad happened, application should be stopped")
	logger.Success("Hooray!! Application ran successfully")

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
