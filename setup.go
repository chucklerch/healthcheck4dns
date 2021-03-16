package main

// Setup the program's configuration, from the config file
//  and evironment and store the results in global variables.

import (
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

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
