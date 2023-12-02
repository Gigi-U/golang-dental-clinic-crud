package utils

import (
	"time"
)
// ConvertDate is a function that takes a time.Time object and converts it to a formatted string.
// The formatted string represents the date in RFC3339 format without the time zone information.
func ConvertDate(fecha time.Time) (string, error) {
	// Format the date in RFC3339 format without the time zone information
	fechaFormateada := fecha.Format(time.RFC3339)[:len(time.RFC3339)-6]
	// Return the formatted date string and a nil error (indicating success)
	return fechaFormateada, nil
}