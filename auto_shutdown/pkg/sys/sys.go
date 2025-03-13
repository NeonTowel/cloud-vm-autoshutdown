// Package sys provides system monitoring and shutdown functionality for Linux VMs.
// [Scope: SYSTEM MONITORING] [Status: Stable]
// Monitors system load and SSH users to determine when to initiate shutdown.

package sys

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// PURPOSE: Start monitoring system load and SSH users to determine shutdown conditions
// WHY: Automates shutdown process based on system metrics to conserve resources
func StartShutdownMonitor(threshold float64, intervals int, sleepTime int) {
	fmt.Println("Purpose: \tMonitor system load and SSH users on a Linux VM")
	fmt.Printf("Shutdown Criteria:\n")
	fmt.Printf("  - System load 5 min average below %.2f\n", threshold)
	fmt.Printf("  - For %d consecutive intervals\n", intervals)
	fmt.Printf("  - Each interval: %d seconds\n", sleepTime)
	fmt.Printf("  - Maximum duration: %.1f minutes\n", float64(intervals)*float64(sleepTime)/60.0)
	fmt.Println("Additional Checks:")
	fmt.Println("  - No logged-in users")
	fmt.Println("Shutdown Process:")
	fmt.Println("  - 5-minute delay before poweroff")
	fmt.Println("=================================")
	fmt.Println("Initiating monitoring loop...")

	count := 0
	for {
		load := getSystemLoad()
		sshUsers := getSSHUsers()

		printStatus(count, load, threshold, intervals, sshUsers)

		if load >= threshold || sshUsers > 0 {
			count = 0
		} else {
			count++
		}

		if count >= intervals {
			if shouldShutdown() {
				initiateShutdown()
			} else {
				count = 0
			}
		}

		time.Sleep(time.Duration(sleepTime) * time.Second)
	}
}

// PURPOSE: Retrieve the system's 5-minute load average
// WHY: Used to assess current system load for shutdown decision
func getSystemLoad() float64 {
	data, _ := os.ReadFile("/proc/loadavg")
	fields := strings.Fields(string(data))
	load, _ := strconv.ParseFloat(fields[2], 64)
	return load
}

// PURPOSE: Count active SSH connections
// WHY: Determines if users are connected via SSH, affecting shutdown decision
func getSSHUsers() int {
	out, _ := exec.Command("ss", "-H", "sport = :22").Output()
	return len(strings.Split(string(out), "\n")) - 1
}

// PURPOSE: Count logged-in users
// WHY: Determines if users are logged in, affecting shutdown decision
func getLoggedInUsers() int {
	out, _ := exec.Command("who").Output()
	return len(strings.Split(string(out), "\n")) - 1
}

// PURPOSE: Initiate system shutdown with a delay
// WHY: Provides a grace period before shutdown to allow for final checks
func initiateShutdown() {
	fmt.Println("Shutting down in 300 seconds ...")
	time.Sleep(300 * time.Second)
	exec.Command("sudo", "poweroff").Run()
}

// PURPOSE: Print current monitoring status
// WHY: Provides real-time feedback on system load and user connections
func printStatus(count int, load float64, threshold float64, intervals int, sshUsers int) {
	var loadString string
	if load >= threshold {
		loadString = "active"
	} else {
		loadString = "idle"
	}
	fmt.Printf("Interval count: %d/%d | System load (5min avg): %.3f (threshold: %.3f, %s) | SSH users: %d | Logged-in users: %d\n", count, intervals, load, threshold, loadString, sshUsers, getLoggedInUsers())
}

// PURPOSE: Determine if conditions are met for shutdown
// WHY: Ensures no users are logged in or connected via SSH before shutdown
func shouldShutdown() bool {
	return getLoggedInUsers() == 0 && getSSHUsers() == 0
}
