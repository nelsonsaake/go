package kernel

import (
	"log"

	"github.com/spf13/cobra"
)

func Start() error {

	// create root command
	var lts = &cobra.Command{Use: "ns"}

	// add routes
	loadRegister(lts)

	// execute command
	err := lts.Execute()
	if err != nil {
		log.Fatalf("error executing command: %v", err)
	}

	return err
}
