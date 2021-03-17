package main

import (
	"crypto/tls"
	"net/http"
	"time"
)

func checkSite(url string) (bool, error) {
	t := time.Duration(5 * time.Second)
	// Create a transport for better control.
	tr := &http.Transport{
		DisableKeepAlives:   true,
		TLSHandshakeTimeout: time.Duration(5 * time.Second),
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: t}

	ticker := time.NewTicker(frequency)
	// Loop
	for {
		// Wait for the ticker to fire.
		<-ticker.C
		// Make the request.
		resp, err := client.Get(url)
		if err != nil {
			results <- false
		}
		resp.Body.Close()

		// Only if the status is 200, success.
		// TODO:  Change to accept more status codes.
		if resp.StatusCode == 200 {
			results <- true
		} else {
			results <- false
		}
	}
}
