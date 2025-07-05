package utils

import (
	"fmt"
	"log"
	"os"
)

// LogInfo logs informational messages to the console.
func LogInfo(message string) {
	log.Println("INFO:", message)
}

// LogError logs error messages to the console and exits the program.
func LogError(message string) {
	log.Println("ERROR:", message)
	os.Exit(1)
}

// FormatCurrency formats a float64 value as a string with two decimal places.
func FormatCurrency(amount float64) string {
	return fmt.Sprintf("%.2f", amount)
}
