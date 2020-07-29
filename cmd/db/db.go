package db

import (
	"os"

	"github.com/spf13/cobra"
)

// DBCmd represents the db command
var DBCmd = &cobra.Command{
	Use:   "db",
	Short: "Shows the db command.",
	Long:  `Shows the db command.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}
