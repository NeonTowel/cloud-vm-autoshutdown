// Package universal manages the auto shutdown process by monitoring system metrics.
// [Scope: SYSTEM MONITORING] [Status: Stable]
// Provides functionality to delay startup and initiate shutdown based on thresholds.

package universal

import (
	"fmt"
	"time"

	"auto_shutdown/pkg/sys"
)

const (
	defaultThreshold    = 0.15
	defaultIntervals    = 15
	defaultSleepTime    = 30
	defaultInitialDelay = 3600 // Default initial delay in seconds (1 hour)
)

// PURPOSE: Handle initial delay before starting the monitoring process
// WHY: Allows for a grace period before system monitoring begins
func handleInitialDelay(initialDelay int) {
	if initialDelay >= 1 {
		fmt.Printf("Waiting for %s before starting...\n", formatTime(initialDelay))
		time.Sleep(time.Duration(initialDelay) * time.Second)
	}
}

// PURPOSE: Monitor system metrics and initiate shutdown if thresholds are exceeded
// WHY: Automates shutdown to conserve resources based on predefined conditions
func MonitorAndShutdown() {
	threshold := getEnvOrDefault("SHUTDOWN_THRESHOLD", defaultThreshold)
	intervals := getEnvOrDefaultInt("SHUTDOWN_INTERVALS", defaultIntervals)
	sleepTime := getEnvOrDefaultInt("SHUTDOWN_SLEEP_TIME", defaultSleepTime)
	initialDelay := getEnvOrDefaultInt("INITIAL_DELAY", defaultInitialDelay) // Get initial delay

	fmt.Println("=== Universal Auto Shutdown ===")
	handleInitialDelay(initialDelay)

	sys.StartShutdownMonitor(threshold, intervals, sleepTime)
}
