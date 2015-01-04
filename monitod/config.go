package monitod

import (
	"path/filepath"
	"os"
	"log"
)

const (
	statdir = "/proc/stat"
)

var (
	cwd string = getDir()
)

func getDir() (dir string) {
	//get the absolutepath of this file
	//os.Args[0] is the filename
	//NOTE: must be installed and not ran with go run
	//go run does it from a tmp dir
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return
}

