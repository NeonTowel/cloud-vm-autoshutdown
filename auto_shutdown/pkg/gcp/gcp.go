package gcp

import (
	"fmt"
	"net/http"

	"auto_shutdown/pkg/sys"
)

const (
	threshold   = 0.15
	intervals   = 15
	sleepTime   = 30
	metadataURL = "http://metadata.google.internal/computeMetadata/v1"
)

func MonitorAndShutdown() {
	if !isGCEVM() {
		fmt.Println("This script only works on GCE VMs, aborting ...")
		return
	}

	fmt.Println("=== GCE VM Auto Shutdown ===")
	sys.StartShutdownMonitor(threshold, intervals, sleepTime)
}

func isGCEVM() bool {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", metadataURL, nil)
	req.Header.Add("Metadata-Flavor", "Google")
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.Header.Get("Metadata-Flavor") == "Google"
}
