package main


import (
	"log"
	
	"github.com/PI-Victor/monito/monitod" //import the daemon pkg
)



func main() {
	if err:= monitod.GetSeries(); err != nil {
		log.Fatal(err)
	}
}
