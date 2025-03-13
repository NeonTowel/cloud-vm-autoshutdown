// Package main is the entry point for the Universal Auto Shutdown application.
// [Scope: APPLICATION ENTRY] [Status: Stable]
// Initiates the monitoring and shutdown process by calling the universal package.

package main

import "auto_shutdown/pkg/universal"

// PURPOSE: Entry point for the application
// WHY: Starts the monitoring and shutdown process
func main() {
	universal.MonitorAndShutdown()
}
