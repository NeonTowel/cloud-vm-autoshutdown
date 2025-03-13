package universal

import (
	"fmt"

	"auto_shutdown/pkg/sys"
)

const (
	threshold = 0.15
	intervals = 15
	sleepTime = 30
)

func MonitorAndShutdown() {
	fmt.Println("=== Universal Auto Shutdown ===")
	sys.StartShutdownMonitor(threshold, intervals, sleepTime)
}
