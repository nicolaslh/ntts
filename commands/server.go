package commands

import (
	"fmt"

	"github.com/nicolaslh/ntts/pkg/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "ntts server start",
	Long:  "ntts server start",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start server")
		server.NewWebsocket()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
