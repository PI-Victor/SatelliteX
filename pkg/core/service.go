package monito

import (
	"fmt"
	"log"

	"github.com/PI-Victor/monito/pkg"
)

// MainService - Main monitoring service core
type MainService struct {
	ConfigFile string
}

func (m *MainService) loadService(mainFilePath string) {
	workdir, err := util.GetWorkDir(mainFilePath)
	if err != nil {
		log.Fatal("An error occured while getting the working directory")
	}
	fmt.Printf("%q", workdir)
}

// Start - Starts the main service
func (m *MainService) Start() {
	m.loadService("String")

}

//CheckConfig - validates the configuration file
func (m *MainService) CheckConfig() {
	fmt.Printf("Checking configuration, %s", m.ConfigFile)
}

//LoadAssets - load and test configured components for runtime
func (m *MainService) LoadAssets() {

}
