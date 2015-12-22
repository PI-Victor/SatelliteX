package config

import (
	"fmt"

	_ "github.com/op/go-logging" // leave the logging package
	_ "gopkg.in/yaml.v2"         // and the yaml package so that i don't forget where they are
)

func loadConfig() {
	fmt.Println("Loading Configuration")
}
