// Package universal provides utility functions for formatting time and retrieving environment variables.
// [Scope: UTILITIES] [Status: Stable]
// Includes functions for converting seconds to human-readable time and fetching environment variables with defaults.

package universal

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PURPOSE: Format seconds into a human-readable time string
// WHY: Provides a user-friendly display of time durations
func formatTime(seconds int) string {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	seconds = seconds % 60

	var parts []string
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%d hour", hours))
		if hours > 1 {
			parts[len(parts)-1] += "s"
		}
	}
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%d minute", minutes))
		if minutes > 1 {
			parts[len(parts)-1] += "s"
		}
	}
	if seconds > 0 {
		parts = append(parts, fmt.Sprintf("%d second", seconds))
		if seconds > 1 {
			parts[len(parts)-1] += "s"
		}
	}

	if len(parts) > 1 {
		return strings.Join(parts[:len(parts)-1], ", ") + " and " + parts[len(parts)-1]
	}
	return strings.Join(parts, ", ")
}

// PURPOSE: Retrieve environment variable as float or use default
// WHY: Allows configuration via environment with fallback
func getEnvOrDefault(envVar string, defaultValue float64) float64 {
	if value, exists := os.LookupEnv(envVar); exists {
		if parsedValue, err := strconv.ParseFloat(value, 64); err == nil {
			return parsedValue
		}
	}
	return defaultValue
}

// PURPOSE: Retrieve environment variable as integer or use default
// WHY: Allows configuration via environment with fallback
func getEnvOrDefaultInt(envVar string, defaultValue int) int {
	if value, exists := os.LookupEnv(envVar); exists {
		if parsedValue, err := strconv.Atoi(value); err == nil {
			return parsedValue
		}
	}
	return defaultValue
}
