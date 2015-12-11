package cmd

import (
	"github.com/PI-Victor/monito/pkg/core"
	"github.com/spf13/cobra"
)

var (
	confFile string
)

var StartCommand = &cobra.Command{
	Use:   "start",
	Short: "Start the server monitor",
	Long:  "Start Monito, the server monitor",
	Run: func(cmd *cobra.Command, args []string) {
		monitoService := &service.MonitoService{}
		monitoService.Start()
	},
}

var CheckConfig = &cobra.Command{
	Use:     "check",
	Short:   "Run a dry-check on the provided config file",
	Long:    "Don't start the monitoring server, just run a check on the config file provided by --config=",
	Example: `monito check --conf=/opt/monito/monito.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		monitoService := &service.MonitoService{ConfigFile: confFile}
		monitoService.CheckConfig()
	},
}

var LoadAssets = &cobra.Command{
	Use:     "load-assets",
	Short:   "Load and test configured assets",
	Long:    "Load and test configured assets",
	Example: `monito load-assets`,
	Run: func(cmd *cobra.Command, args []string) {
		monitoService := &service.MonitoService{ConfigFile: confFile}
		monitoService.LoadAssets()
	},
}

func init() {
	StartCommand.PersistentFlags().StringVar(&confFile, "config", "c", "Specify a yaml configuration file")
	CheckConfig.PersistentFlags().StringVar(&confFile, "config", "c", "Specify a yaml configuration file")
	LoadAssets.PersistentFlags().StringVar(&confFile, "config", "c", "Specify a yaml configuration file")
}
