package monitod


import (
	"fmt"
	"log"
	_ "io" //silence the unused imports with the blank operator
	linuxproc "github.com/c9s/goprocinfo/linux"
	_ "github.com/sevlyar/go-daemon"
)

type CpuStat struct {
	user      float64
	nice      float64
	system    float64
	idle      float64
	iowait    float64
	
}

func (c *CpuStat) getCpuStat() {
	
}

func GetSeries() {
	//just print the dir for the file for now
	fmt.Println("we are in the dir", cwd)
	fmt.Println("This is a test")
	stat, err := linuxproc.ReadStat(statdir)
	
	if err != nil {
		log.Fatal(err)
	}
	
	for _, s := range stat.CPUStats {
		fmt.Println(
			"\nCPU Stats\n", 
			"\nUser: ", s.User,
			"\nNice: ", s.Nice,
			"\nSystem: ", s.System,
			"\nIdle: ", s.Idle,
			"\nIoWait: ", s.IOWait,
		)
	}
	fmt.Println(
		"\nCpu stats combined: ", stat.CPUStatAll, 
		"\nCpu stats individual: ", stat.CPUStats,
		"\nNumber of processes: ", stat.Processes,
		"\nUp since: ", stat.BootTime,
	)
}
