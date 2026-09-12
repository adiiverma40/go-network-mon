package network

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func ReadPortalInfo() error {
	portalURL := "http://172.11.0.1:8090/"
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	start := time.Now()

	resp, err := client.Get(portalURL)
	if err != nil {
		return fmt.Errorf("failed to reach portal: %v", err)
	}
	defer resp.Body.Close()
	latency := time.Since(start)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	fmt.Println("--- Portal info ---")
	fmt.Printf("Latency:        %v\n", latency)
	fmt.Printf("Status Code:    %d\n", resp.StatusCode)

	fmt.Printf("Final URL:      %s\n", resp.Request.URL.String())
	fmt.Println("--------------------------")
	fmt.Printf("Headers Data:   %v\n", resp.Header)
	fmt.Println("--------------------------")
	fmt.Printf("Page Source:\n%s\n", string(body))

	return nil
}
