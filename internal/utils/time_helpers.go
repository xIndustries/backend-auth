package utils

import (
	"time"
)

// FormatTimestamp formats a time.Time object into RFC3339 format.
func FormatTimestamp(t time.Time) string {
	return t.Format(time.RFC3339)
}

// ParseTimestamp parses an RFC3339 timestamp into a time.Time object.
func ParseTimestamp(timestamp string) (time.Time, error) {
	return time.Parse(time.RFC3339, timestamp)
}

// GetCurrentTimestamp returns the current time in RFC3339 format.
func GetCurrentTimestamp() string {
	return time.Now().Format(time.RFC3339)
}
