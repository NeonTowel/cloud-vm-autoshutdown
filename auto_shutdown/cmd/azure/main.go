package main

import "auto_shutdown/pkg/azure"

func main() {
	azure.MonitorAndShutdown()
}
