package utils

import (
	"log"
	"os"
)

var logger *log.Logger

// InitLogger initializes the logger and directs logs to a file.
func InitLogger(logFile string) error {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	logger = log.New(file, "", log.LstdFlags|log.Lshortfile)
	log.SetOutput(file) // Direct global log calls to the same file
	return nil
}

// Info logs informational messages.
func Info(message string) {
	logger.Printf("[INFO] %s\n", message)
}

// Error logs error messages.
func Error(message string) {
	logger.Printf("[ERROR] %s\n", message)
}

// Debug logs debug messages.
func Debug(message string) {
	logger.Printf("[DEBUG] %s\n", message)
}
