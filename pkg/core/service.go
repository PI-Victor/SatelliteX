package service

import (
	"fmt"
	"log"

	"github.com/PI-Victor/monito/pkg"
)

// Main monitoring service core
type MonitoService struct {
	workDir    string
	ConfigFile string
}

func (m *MonitoService) loadService(mainFilePath string) {
	workdir, err := util.GetWorkDir(mainFilePath)
	if err != nil {
		log.Fatal("An error occured while getting the working directory")
	}
	m.workDir = workdir
}

func (m *MonitoService) Start() {
	m.loadService("String")
	fmt.Printf("Starting %s", m.workDir)
}

func (m *MonitoService) CheckConfig() {
	fmt.Printf("Checking configuration, %s", m.ConfigFile)
}

func (m *MonitoService) LoadAssets() {

}
