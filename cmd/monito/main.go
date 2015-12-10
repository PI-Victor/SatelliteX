package main

import (
	"github.com/PI-Victor/monito/pkg/cmd"
	"github.com/spf13/cobra"
)

func main() {
	RootCmd.AddCommand(cmd.StartCommand)
	RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use:   "monito",
	Short: "Server Monitoring Service",
	Long:  "Monito - Server monitoring service",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
