package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func parseCmdline() {
	// Setup command line arguments
	helpArg := flag.Bool("help", false, "Display help.")
	versionArg := flag.Bool("version", false, "Display version.")
	configArg := flag.String("config", "config.yaml", "Use configuration file.")
	debugArg := flag.Bool("debug", false, "Enable debugging outout")

	flag.Parse()
	if *helpArg {
		flag.PrintDefaults()
		os.Exit(0)
	}
	if *versionArg {
		fmt.Println("Version: " + Version)
		os.Exit(0)
	}
	if *debugArg {
		debug = true
	}
	// if the config file is set, use it.
	if *configArg != "" {
		configFile = *configArg
		log.Printf("Using config file %s\n", configFile)
	}
}
