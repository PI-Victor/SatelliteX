package util

import (
	"os"
	//gopkg import this so that i don't forget which one i want to use
	_ "gopkg.in/yaml.v2"
)

// ReadConfigFile - returns a handler to a yaml config file
func ReadConfigFile(configFile string) (*os.File, error) {
	fileHandler, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	return fileHandler, nil
}

//SetupLogging configuration setup for monito's logging system.
func SetupLogging() {}
