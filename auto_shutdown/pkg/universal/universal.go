package universal

import (
	"fmt"
	"os"
	"strconv"

	"auto_shutdown/pkg/sys"
)

const (
	defaultThreshold = 0.15
	defaultIntervals = 15
	defaultSleepTime = 30
)

func getEnvOrDefault(envVar string, defaultValue float64) float64 {
	if value, exists := os.LookupEnv(envVar); exists {
		if parsedValue, err := strconv.ParseFloat(value, 64); err == nil {
			return parsedValue
		}
	}
	return defaultValue
}

func getEnvOrDefaultInt(envVar string, defaultValue int) int {
	if value, exists := os.LookupEnv(envVar); exists {
		if parsedValue, err := strconv.Atoi(value); err == nil {
			return parsedValue
		}
	}
	return defaultValue
}

func MonitorAndShutdown() {
	threshold := getEnvOrDefault("SHUTDOWN_THRESHOLD", defaultThreshold)
	intervals := getEnvOrDefaultInt("SHUTDOWN_INTERVALS", defaultIntervals)
	sleepTime := getEnvOrDefaultInt("SHUTDOWN_SLEEP_TIME", defaultSleepTime)

	fmt.Println("=== Universal Auto Shutdown ===")
	sys.StartShutdownMonitor(threshold, intervals, sleepTime)
}
