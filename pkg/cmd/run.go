package cmd

import (
	"github.com/PI-Victor/monito/pkg/core"
	"github.com/spf13/cobra"
)

var (
	confFile string
)

// StartCommand starts the main server
var StartCommand = &cobra.Command{
	Use:   "start",
	Short: "Start the server monitor",
	Long:  "Start Monito, the server monitor",
	Run: func(cmd *cobra.Command, args []string) {
		monitoService := &monito.MainService{
			ConfigFile: confFile,
		}
		monitoService.Start()
	},
}

func init() {
	StartCommand.PersistentFlags().StringVar(&confFile, "config", "c", "Specify a configuration file")
}
