package monitod

import (
	"fmt"
)

type dataSet struct {
	f float64
}

func (s dataSet) printDataSet() {
	fmt.Println(s.f)
}