package server

import (
	"github.com/muharihar/d3ta-go/interface/restapi"
	"github.com/spf13/cobra"
)

// restAPICmd represents the restapi server command
var restAPICmd = &cobra.Command{
	Use:   "restapi",
	Short: "Shows the restapi server command.",
	Long:  `Shows the restapi server command.`,
	Run: func(cmd *cobra.Command, args []string) {

		restapi.StartRestAPIServer()
	},
}

func init() {
	ServerCmd.AddCommand(restAPICmd)
}
