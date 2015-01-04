package monito


import (
	_ "github.com/c9s/goprocinfo/linux"
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
