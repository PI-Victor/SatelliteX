package main

import (
	"log"

	"github.com/pi-victor/monito/monitod"
)

//import the daemon pkg

func main() {
	if err := monitod.GetSeries(); err != nil {
		log.Fatal(err)
	}

}
