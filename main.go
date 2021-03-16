package main

import (
	"fmt"
	"log"
	"time"
)

// Global variables.
var (
	url          = "http://google.com/"
	frequency, _ = time.ParseDuration("60s")
	logfile      = "out.log"
)

// Config file structure.
type Conf struct {
	HealthCheck struct {
		Server     string
		Port       int
		IgnoreCert bool
		Frequency  string
	}
	Publish struct {
		Log struct {
			Filename string
		}
		Cloudwatch struct {
			Name string
		}
	}
}

// Main
func main() {
	conf, err := readConf("config.yaml")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Printf("Dump: %+v\n", conf)
	setConf(conf)
	results, err := checkSite(url)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	if results {
		fmt.Println("Request good.")
	}

}
