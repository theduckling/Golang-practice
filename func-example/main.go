package main

import (
	"fmt"
	"time"
)

func ConvertToDigitalFormat(hour, minute, second int) string {
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func ExtractTimeUnits(seconds int) (int, int, int) {
	duration := time.Duration(seconds) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds = seconds % 60
	// return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	return hours, minutes, seconds
}
