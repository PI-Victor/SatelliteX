package monito

import (
	"fmt"
	"sync"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/docker"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

// ReadSystemMetrics - log psutil local system metrics.
// TODO: implement concurrency
func ReadSystemMetrics() {

	var wg sync.WaitGroup
	wg.Add(3)

	// performance
	go getSystemMemory(&wg)
	go getSystemSwap(&wg)
	go getSystemLoad(&wg)
	wg.Wait()
}

// ReadSystemInfo read one time needed system information
func ReadSystemInfo() {
	go hostInfo()
	go cpuInfo()

	// other
	go dockerInfo()
}

func getSystemMemory(wg *sync.WaitGroup) {
	defer wg.Done()
	vmem, _ := mem.VirtualMemory()
	fmt.Printf("This is the memory info: %v \n", vmem)
}

func getSystemSwap(wg *sync.WaitGroup) {
	defer wg.Done()
	swap, _ := mem.SwapMemory()
	fmt.Printf("This is the swap info: %v \n", swap)

}

func getSystemLoad(wg *sync.WaitGroup) {
	defer wg.Done()
	loadInfo, _ := load.LoadAvg()
	fmt.Printf("Load average info: %v \n", loadInfo)
}

// read this only on config load and at config reload, otherwise remove from pool.
func cpuInfo() {
	cpu, _ := cpu.CPUInfo()
	fmt.Printf("This is the CPU info: %v \n", cpu)
}

// slaves shouldn't run their own api, but rather report to master
// this is where to get the hostname
func hostInfo() *host.HostInfoStat {
	hostInfo, _ := host.HostInfo()
	fmt.Printf("This is the host info: %+v \n", hostInfo)
	return hostInfo
}

// skip on pool
func dockerInfo() {
	dockerIDList, _ := docker.GetDockerIDList()
	fmt.Println(dockerIDList)
	// these two function need a valid docker container as a parameter
	// https://godoc.org/github.com/shirou/gopsutil/docker
	cgroupCPU, _ := docker.CgroupCPU("test", "test")
	cgroupMem, _ := docker.CgroupMem("test", "test")

	fmt.Printf("Docker list info: %q \n, %q \n, %v \n", dockerIDList, cgroupCPU, cgroupMem)
}
