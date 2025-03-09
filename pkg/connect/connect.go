package connect

import (
	"net/http"
	"time"
)

var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: 2 * time.Second,
}

// Ping check if url is correct can be reached
func Ping(url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	resp.Body.Close()

	// only get with 200 is ok
	return resp.StatusCode == http.StatusOK
}
