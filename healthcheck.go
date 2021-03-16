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
	resp, err := client.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return true, nil
	} else {
		return false, err
	}
}
