package main

import "testing"

func TestCheckSite(t *testing.T) {
	url := "http://www.google.com/"
	a, err := checkSite(url)
	if err != nil && !a {
		t.Errorf("checkSite on %v failed, error %v", url, err)
	}
	// url = "http://bad.site/"
	// a, err = checkSite(url)
	// if err != nil || a {
	// 	t.Errorf("checkSite on %v failed, error %v", url, err)
	// }
}
