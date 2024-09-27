package logger

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var Logger *log.Logger

// func init() {

// 	logFile, err := os.OpenFile("duffman.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatalf("failed to open log file: %v", err)
// 	}

// 	Logger = log.NewWithOptions(logFile, log.Options{
// 		ReportCaller:    true,
// 		ReportTimestamp: true,
// 		TimeFormat:      time.Kitchen,
// 		// Prefix:          "Baking 🍪 ",
// 	})
// 	Logger.SetLevel(log.InfoLevel)
// 	Logger.SetTimeFormat("2006-01-02 15:04:05")

// }

func Init() {

	logFile, err := os.OpenFile("duffman.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	Logger = log.NewWithOptions(logFile, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		// Prefix:          "Baking 🍪 ",
	})
	Logger.SetLevel(log.DebugLevel)
	Logger.SetTimeFormat("2006-01-02 15:04:05")

}
