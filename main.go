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

	// Get the results of a healthcheck
	results, err := checkSite(url)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	if results {
		log.Println("Success.")
	}
}
