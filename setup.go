package main

// Setup the program's configuration, from the config file
//  and evironment and store the results in global variables.

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"
)

// Config file structure.
type Conf struct {
	HealthCheck struct {
		Server     string
		Port       int
		Path       string
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
	url = "http://" + c.HealthCheck.Server + ":" + strconv.Itoa(c.HealthCheck.Port) + c.HealthCheck.Path
	//logfile = c.Publish.Log.Filename
	frequency, _ = time.ParseDuration(c.HealthCheck.Frequency)
	if debug {
		log.Printf("URL: %s\n", url)
		log.Printf("FREQ: %s\n", frequency.String())
		log.Printf("IGNORE CERT: %t\n", c.HealthCheck.IgnoreCert)
	}
}
