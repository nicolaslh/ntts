package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ntts",
	Short: "Notify the transit station",
	Long:  "Notify the transit station",
}

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	return err
}
