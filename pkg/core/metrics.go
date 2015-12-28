package monito

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/docker"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

// ReadSystemMetrics - log psutil local system metrics.
// TODO: implement async
func ReadSystemMetrics() {
	// system information
	getHost()
	cpuInfo()

	// performance
	readMemory()
	swapInfo()
	loadInfo()

	// other
	dockerInfo()
}

// slaves shouldn't run their own api, but rather report to master
// this is where to get the hostname

func getHost() {
	hostInfo, _ := host.HostInfo()
	fmt.Printf("This is the host info: %+v \n", hostInfo)
}

func readMemory() {
	vmem, _ := mem.VirtualMemory()
	fmt.Printf("This is the memory info: %v \n", vmem)
}

func swapInfo() {
	swap, _ := mem.SwapMemory()
	fmt.Printf("This is the swap info: %v \n", swap)
}

func cpuInfo() {
	cpu, _ := cpu.CPUInfo()
	fmt.Printf("This is the CPU info: %v \n", cpu)
}

func dockerInfo() {
	dockerIDList, _ := docker.GetDockerIDList()
	fmt.Println(dockerIDList)
	// these two function need a valid docker container as a parameter
	// https://godoc.org/github.com/shirou/gopsutil/docker
	cgroupCPU, _ := docker.CgroupCPU("test", "test")
	cgroupMem, _ := docker.CgroupMem("test", "test")

	fmt.Printf("Docker list info: %v \n, %v \n, %v \n", dockerIDList, cgroupCPU, cgroupMem)
}

func loadInfo() {
	loadInfo, _ := load.LoadAvg()
	fmt.Printf("Load average info: %v \n", loadInfo)
}
