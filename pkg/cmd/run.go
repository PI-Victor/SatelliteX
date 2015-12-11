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

func init() {
	StartCommand.PersistentFlags().StringVar(&confFile, "config", "c", "Specify a yaml configuration file")
}
