package main


import (
	_ "log"
	"fmt"
	daemon "github.com/sevlyar/go-daemon"
	_ "github.com/PI-Victor/monito/monitod" //import the daemon pkg
)



func main() {


}


// if err:= monitod.GetSeries(); err != nil {
// 	log.Fatal(err)
// }