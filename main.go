package main

import (
	"log"
	"time"
)

//
const Version = "v0.0.1"

// Global variables.
var (
	debug      bool
	url        string
	frequency  time.Duration
	configFile string
)
var results = make(chan bool)

// Main
func main() {
	// Startup.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("Sarting version %s.\n", Version)

	// Parse command line.
	parseCmdline()

	// Read the config file.
	if configFile != "" {
		conf, err := readConf(configFile)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		// Dump the config contents.
		if debug {
			log.Printf("Dump: %+v\n", conf)
		}
		// Set variables according to the config file.
		setConf(conf)
	}

	// Start the health checking thread.
	log.Println("Starting checkSite thread.")
	go checkSite(url)

	// Loop of Infinity.
	for {
		// Wait for results to return.
		healthy, ok := <-results
		if !ok && !healthy {
			log.Println("Health check failure.")
		}
		if healthy {
			log.Println("Health check success.")
		}
	}
}
