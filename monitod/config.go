package utils

import (
	"fmt"
	"os"
	"errors"
	"log"
	"io/ioutil"
	"path/filepath"
)

var (
	WorkDir = getWorkDir()
	configDirEmpty =  errors.New("Application directory is empty")
)

const (
	AppDir = "/etc/opt/monito/"
	ConfigDir = "/etc/opt/dataaps/config1/"
)

type ConfigSet struct {
	recursiveSearch bool
	filterConfig string
	ymlFiles []string
}

func (confSet *ConfigSet) getConfigFiles(dirPath string) (ymlFiles []string, err error) {
	// returns a sorted list of yml config files
	// from a given path
	// if the search is recursive, it searches the directories inside
	// if it returns empty, it panics, nothing to build on to from here
	files, err := ioutil.ReadDir(dirPath)

	if err != nil {
		return err
	}

	for _, file := range files {
		configFile := filepath.Join(dirPath, file.Name())
		if fileExt := filepath.Ext(file.Name()); fileExt == ".yml" {
			confSet.ymlFiles = append(confSet.ymlFiles, configFile)
		}
		if file.IsDir() && confSet.recursiveSearch {
			confSet.getConfigFiles(configFile)
		}
	}

	if arrayLen := len(confSet.ymlFiles); arrayLen == 0 {
		return ymFiles, configDirEmpty
	}

	return confSet.ymlFiles, nil
}

func (confSet *ConfigSet) ParseConfigFiles() {
	//temp for testing
	confSet.recursiveSearch = true
	configFiles, err := confSet.getConfigFiles(ConfigDir)

	if err != nil {
		log.Fatal(err)
	}

	for _, filePath := range configFiles {
		fmt.Println(filePath)
	}
}

func getWorkDir() (dir string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		log.Fatal(err)
	}
	return

}
