package sys

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Public functions

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

// Private functions

func getSystemLoad() float64 {
	data, _ := os.ReadFile("/proc/loadavg")
	fields := strings.Fields(string(data))
	load, _ := strconv.ParseFloat(fields[2], 64)
	return load
}

func getSSHUsers() int {
	out, _ := exec.Command("ss", "-H", "sport = :22").Output()
	return len(strings.Split(string(out), "\n")) - 1
}

func getLoggedInUsers() int {
	out, _ := exec.Command("who").Output()
	return len(strings.Split(string(out), "\n")) - 1
}

func initiateShutdown() {
	fmt.Println("Shutting down in 300 seconds ...")
	time.Sleep(300 * time.Second)
	exec.Command("sudo", "poweroff").Run()
}

func printStatus(count int, load float64, threshold float64, intervals int, sshUsers int) {
	var loadString string
	if load >= threshold {
		loadString = "active"
	} else {
		loadString = "idle"
	}
	fmt.Printf("Interval count: %d/%d | System load (5min avg): %.3f (threshold: %.3f, %s) | SSH users: %d | Logged-in users: %d\n", count, intervals, load, threshold, loadString, sshUsers, getLoggedInUsers())
}

func shouldShutdown() bool {
	return getLoggedInUsers() == 0 && getSSHUsers() == 0
}
