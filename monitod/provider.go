package monitod

import (
	"fmt"
)

type dataSet struct {
	 name string
	 value float64
	 type []string{
	 	"float",
	 	"string"
	 	"int"
	 	"gauge"
	 }
}

func (s dataSet) printDataSet() {
	fmt.Println(s.f)
}