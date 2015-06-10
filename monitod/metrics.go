package monitod

import (
	"fmt"
	_ "log"
	_ "io" 
	linuxproc "github.com/c9s/goprocinfo/linux"
)


type individualCpuStats struct {
	user      uint64
	nice      uint64
	system    uint64
	idle      uint64
	iowait    uint64
}

func(c *individualCpuStats) filterCpuStat(metricFilter string) {

	if len(metricFilter) != 0 {
		//len() is faster than string comparison
		//this means we want to filter the returned
		//array based on what we got
		fmt.Println("test")
	}

}

//remember to return a map with the stats collected
func metricsProvider()  (err error) {

	cpuStat, err := linuxproc.ReadStat("/proc/stat")

	if err != nil {
		return err
	}

	for _, cpuStat := range cpuStat.CPUStats {
		iCpuStats := individualCpuStats {
			cpuStat.User,
			cpuStat.Nice,
			cpuStat.System,
			cpuStat.Idle,
			cpuStat.IOWait,
		}
		fmt.Println(iCpuStats)
	}
	return nil
}

func GetSeries() (err error) {

	//just print the dir for the file for now
	//fmt.Println("we are in the dir", cwd)
	fmt.Println("This is a test")

	err = metricsProvider()

	if err != nil {
		return err
	}

	// fmt.Println(
	// 	"\nCpu stats combined: ", stat.CPUStatAll, 
	// 	"\nCpu stats individual: ", stat.CPUStats,
	// 	"\nNumber of processes: ", stat.Processes,
	// 	"\nUp since: ", stat.BootTime,
	// )
	return nil
}
