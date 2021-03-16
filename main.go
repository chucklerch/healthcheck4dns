package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
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

// Read the config file and parse the yaml.
func readConf(filename string) (*Conf, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &Conf{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}

func setConf(c *Conf) {
	logfile = c.Publish.Log.Filename
	frequency, _ = time.ParseDuration(c.HealthCheck.Frequency)
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
