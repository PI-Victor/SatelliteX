package main

import (
	"fmt"
	"os"
	"log"
	_ "io" //silence the unused imports with the blank operator
	"path/filepath"
	linuxproc "github.com/c9s/goprocinfo/linux"
	_ "github.com/sevlyar/go-daemon"
)


func main() {
	//just print the dir for the file for now
	cwd := getdir()
	fmt.Println("we are in the dir", cwd)
	stat, err := linuxproc.ReadStat("/proc/stat")
	
	if err != nil {
		log.Fatal(err)
	}
	
	for _, s := range stat.CPUStats {
		fmt.Println(
			s.User,
			s.Nice,
			s.System,
			s.Idle,
			s.IOWait,
		)
	}
	fmt.Println(
		"Cpu all stats", stat.CPUStatAll, 
		stat.CPUStats,
		stat.Processes,
		stat.BootTime,
	)
	
}


func getdir() (dir string) {
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


