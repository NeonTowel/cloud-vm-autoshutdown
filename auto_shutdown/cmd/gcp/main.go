package main

import "auto_shutdown/pkg/gcp"

func main() {
	gcp.MonitorAndShutdown()
}
