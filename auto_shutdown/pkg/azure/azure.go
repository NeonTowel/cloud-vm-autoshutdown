package azure

import (
	"fmt"
	"net/http"

	"auto_shutdown/pkg/sys"
)

const (
	threshold   = 0.15
	intervals   = 15
	sleepTime   = 30
	metadataURL = "http://169.254.169.254/metadata/instance"
)

func MonitorAndShutdown() {
	if !isAzureVM() {
		fmt.Println("This script only works on Azure VMs, aborting ...")
		return
	}

	fmt.Println("=== Azure VM Auto Shutdown ===")
	sys.StartShutdownMonitor(threshold, intervals, sleepTime)
}

func isAzureVM() bool {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", metadataURL+"?api-version=2021-01-01", nil)
	req.Header.Add("Metadata", "true")
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
