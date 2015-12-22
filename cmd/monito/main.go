package main

import (
	"github.com/PI-Victor/monito/pkg/cmd"
	"github.com/spf13/cobra"
)

func main() {
	RootCmd.AddCommand(cmd.StartCommand)
	RootCmd.AddCommand(cmd.CheckConfig)
	RootCmd.AddCommand(cmd.LoadAssets)
	RootCmd.Execute()
}

// RootCmd main command for monito
var RootCmd = &cobra.Command{
	Use:   "monito",
	Short: "Server Monitoring Service",
	Long:  "Monito - A resilient system monitoring service",
	Example: `monito start --config=/opt/monito/monito.yaml
monito check --config=/opt/monito/monito.yaml
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
