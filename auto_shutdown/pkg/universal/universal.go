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
	silenceWarning      = false
)

// PURPOSE: Handle initial delay before starting the monitoring process
// WHY: Allows for a grace period before system monitoring begins
func handleInitialDelay(initialDelay int) {
	if initialDelay >= 1 {
		fmt.Printf("Waiting for %s before starting...\n", formatTime(initialDelay))
		time.Sleep(time.Duration(initialDelay) * time.Second)
	}
}

func preflightChecks() {
	silenceWarning := getEnvOrDefaultBool("AUTO_SHUTDOWN_SILENCE_AZURE_WARNING", false)

	if !silenceWarning && sys.IsAzureVM() {
		fmt.Println("\nAzure VM detected, action required!")
		fmt.Print("=====================================\n\n")
		fmt.Print("This VM will Power Off by Auto Shutdown but not be deallocated!\n(There's a tool for that :P)\n\n")
		fmt.Print("This VM may incur costs unless properly deallocated.\n\n")
		fmt.Println("To fully deallocate the VM, you need to:")
		fmt.Println("1. Assign MSI to the VM")
		fmt.Println("2. Assign permissions for the MSI to the VM:")
		fmt.Println("   - Microsoft.Compute/virtualMachines/deallocate/action")
		fmt.Println("   - Microsoft.Compute/virtualMachines/start/action")
		fmt.Println("   - Microsoft.Compute/virtualMachines/stop/action")
		fmt.Println("   - Microsoft.Compute/virtualMachines/restart/action")
		fmt.Println("   - Microsoft.Compute/virtualMachines/powerOff/action")
		fmt.Println("   - Microsoft.Compute/virtualMachines/delete/action")
		fmt.Println("   - or use 'Virtual Machine Contributor' role")
		fmt.Print("3. Install the Azure VM Deallocate SystemD service to deallocate the VM on shutdown.\n\n")
		fmt.Print("We provide a systemd service to deallocate the VM on shutdown, but you need to ensure the pre-requisites are met.\n\n")
		fmt.Print("Please refer to the following repository for more information:\n")
		fmt.Print("https://github.com/NeonTowel/cloud-vm-autoshutdown\n\n")
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
	preflightChecks()
	handleInitialDelay(initialDelay)

	sys.StartShutdownMonitor(threshold, intervals, sleepTime)
}
