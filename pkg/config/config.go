package config

import (
	"os"
	//gopkg import this so that i don't forget which one i want to use
	_ "gopkg.in/yaml.v2"
)

// Config - main config structure
type Config struct {
	// where to get aditional files for parsing (future)
	workDirs filePaths
	// interval used for pooling
	poolTime int
	// if this instance of the app is a master node.
	isMaster bool
	// true if config should be reloaded
	// send this across the cluster for all slaves to reload?
	isReloaded bool
}

type filePaths map[string]string

// ReadConfigFile - returns a handler to a yaml config file
func ReadConfigFile(configFile string) (*os.File, error) {
	fileHandler, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	return fileHandler, nil
}
