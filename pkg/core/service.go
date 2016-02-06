package monito

import (
	"fmt"
	"time"

	"github.com/PI-Victor/monito/pkg"
	"github.com/PI-Victor/monito/pkg/log"
)

// Server - Main monitoring service core struct
type Server struct {
	configDir  string
	ConfigFile string
	bindAddr   string
	// we will use name as the namespace for our current instance
	// of the server in the API
	name     string
	MainNode bool
}

// New - create a new server instance
func New(confFile string) *Server {
	return &Server{ConfigFile: confFile}
}

func (m *Server) loadServices() error {
	InitializeAPI(m)
	validateConfig(m.ConfigFile)
	return nil
}

// Start - Starts the main service
func (m *Server) Start() {
	log.Info("Loading Services for monito...")
	if err := m.loadServices(); err != nil {
		log.Panic("Failed to start monito... ")
	}
	for {
		log.Info("Logging System metrics...")
		ReadSystemMetrics()
		time.Sleep(10 * time.Second)
	}
}

// ValidateAssets - load and test configured components for runtime
func (m *Server) ValidateAssets() error {
	//InitializeAPI()
	return nil
}

func validateConfig(confFile string) error {
	// if there's no config file passed we generate a default one.
	if len(confFile) == 0 {
		return nil
	}

	configFile, err := util.ReadConfigFile(confFile)
	if err != nil {
		return err
	}
	fmt.Println(configFile)
	return nil
}
