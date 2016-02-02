package cmd

import (
	"github.com/PI-Victor/monito/pkg/core"
	"github.com/spf13/cobra"
)

// ValidateAssets - validates assets from the passed config file
var ValidateAssets = &cobra.Command{
	Use:     "validate-assets",
	Short:   "Validate configured assets",
	Long:    "Validate configured assets",
	Example: `monito validate-assets`,
	Run: func(cmd *cobra.Command, args []string) {
		monitoService := monito.New(confFile)
		monitoService.ValidateAssets()
	},
}

func init() {
	ValidateAssets.PersistentFlags().StringVar(&confFile, "config", "", "Specify a yaml configuration file")
}
