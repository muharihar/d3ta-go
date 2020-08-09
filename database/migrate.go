package database

import (
	"fmt"

	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/initialize"
)

// RunDBMigrate run DBMigrate
func RunDBMigrate() {
	fmt.Println("RunDBMigrate...")

	h, err := handler.NewHandler()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	c, _, err := config.NewConfig("./")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	h.SetConfig(c)

	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	// Migration
	fmt.Println("\nRun::M01FirstInstall...")
	if err := M01FirstInstall(h); err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println("\nDone::M01FirstInstall!")

	fmt.Println("\nDone!")
}
