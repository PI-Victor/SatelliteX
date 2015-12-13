package monito

import (
	"fmt"
	"log"

	"github.com/PI-Victor/monito/pkg"
)

// Main monitoring service core
type Service struct {
	workDir    string
	ConfigFile string
}

func (m *Service) loadService(mainFilePath string) {
	workdir, err := util.GetWorkDir(mainFilePath)
	if err != nil {
		log.Fatal("An error occured while getting the working directory")
	}
	m.workDir = workdir
}

func (m *Service) Start() {
	m.loadService("String")
	fmt.Printf("Starting %s", m.workDir)
}

func (m *Service) CheckConfig() {
	fmt.Printf("Checking configuration, %s", m.ConfigFile)
}

func (m *Service) LoadAssets() {

}
