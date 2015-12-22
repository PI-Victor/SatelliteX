package monito

import "fmt"

// MainService - Main monitoring service core
type MainService struct {
	ConfigFile string
}

func (m *MainService) loadService(mainFilePath string) {

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
