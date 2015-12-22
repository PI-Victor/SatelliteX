package cmd

import (
	"github.com/PI-Victor/monito/pkg/core"
	"github.com/spf13/cobra"
)

// CheckConfig cobra command that validates the configuration
var CheckConfig = &cobra.Command{
	Use:     "check",
	Short:   "Run a dry check on the provided config file",
	Long:    "Don't start the monitoring server, just run a check on the config file provided by --config=",
	Example: `monito check --config=/opt/monito/monito.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		monitoService := &monito.MainService{ConfigFile: confFile}
		monitoService.CheckConfig()
	},
}

// LoadAssets reads the served config file and tries to loads assets
var LoadAssets = &cobra.Command{
	Use:     "load-assets",
	Short:   "Load and test configured assets",
	Long:    "Load and test configured assets",
	Example: `monito load-assets`,
	Run: func(cmd *cobra.Command, args []string) {
		monitoService := &monito.MainService{ConfigFile: confFile}
		monitoService.LoadAssets()
	},
}

func init() {
	CheckConfig.PersistentFlags().StringVar(&confFile, "config", "c", "Specify a yaml configuration file")
	LoadAssets.PersistentFlags().StringVar(&confFile, "config", "c", "Specify a yaml configuration file")
}
