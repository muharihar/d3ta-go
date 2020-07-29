package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// D3TAgoVersion variable
var D3TAgoVersion = "0.1.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the version of d3ta-go.",
	Long:  "Shows the version of d3ta-go.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("d3ta-go version ", D3TAgoVersion)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
