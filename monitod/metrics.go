package monitod


import (
	"fmt"
	"log"
	_ "io" //silence the unused imports with the blank operator
	linuxproc "github.com/c9s/goprocinfo/linux"
	_ "github.com/sevlyar/go-daemon"
)

type individualCpuStats struct {
	user      uint64
	nice      uint64
	system    uint64
	idle      uint64
	iowait    uint64
}

func(c *individualCpuStats) filterCpuStat(metricFilter string) {
	if len(mectriFilter) != 0 {
		//this means we want to filter the returned
		//array based on what we got
		
	}
	return
}

func metricsProvider()  {
	stat, err := linuxproc.ReadStat(statdir)
	if err != nil {
		log.Fatal(err)
	}
	for _, cpuStat := range stat.CPUStats {
		iCpuStats := individualCpuStats{
			cpuStat.User,
			cpuStat.Nice,
			cpuStat.System,
			cpuStat.Idle,
			cpuStat.IOWait,
		}
		fmt.Println(iCpuStats)
	}
}

func GetSeries() {
	//just print the dir for the file for now
	fmt.Println("we are in the dir", cwd)
	fmt.Println("This is a test")
	stat, err := linuxproc.ReadStat(statdir)
	
	if err != nil {
		log.Fatal(err)
	}

	metricsProvider()

	fmt.Println(
		"\nCpu stats combined: ", stat.CPUStatAll, 
		"\nCpu stats individual: ", stat.CPUStats,
		"\nNumber of processes: ", stat.Processes,
		"\nUp since: ", stat.BootTime,
	)
}
