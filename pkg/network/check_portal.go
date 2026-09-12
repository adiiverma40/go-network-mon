package network

import (
	"fmt"
	"net/http"
	"time"
)
var url = "http://clients3.google.com/generate_204"

func CheckPortal() (bool , error){
	client := &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {return http.ErrUseLastResponse},
	}

	res , err := client.Get(url)

	if err != nil {
		return false, fmt.Errorf("network error (disconnected or DNS failure): %v", err)
		
	}
	defer res.Body.Close()

	switch {
		case res.StatusCode == http.StatusNoContent: // 204
			return false, nil
		case res.StatusCode >= 300 && res.StatusCode < 400: // 30x Redirect
			portalURL := res.Header.Get("Location")
			if portalURL != "" {
				fmt.Printf("Intercepted! Captive portal URL is: %s\n", portalURL)
			} else {
				fmt.Println("Intercepted, but no Location header was provided by the portal.")
			}
			return true, nil
		case res.StatusCode == http.StatusOK: // 200
			return true, nil
		default:
			return true, fmt.Errorf("unexpected status code: %d", res.StatusCode)
		}	
}