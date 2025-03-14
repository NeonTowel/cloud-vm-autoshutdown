package sys

import "net/http"

func IsAzureVM() bool {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://169.254.169.254/metadata/instance?api-version=2021-01-01", nil)
	req.Header.Add("Metadata", "true")
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
