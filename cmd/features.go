package cmd

import (
	"github.com/muharihar/d3ta-go/cmd/db"
	"github.com/muharihar/d3ta-go/cmd/server"
)

func init() {
	RootCmd.AddCommand(db.DBCmd)
	RootCmd.AddCommand(server.ServerCmd)
}
