package cmd

import (
	"github.com/AmadlaOrg/hery-playground/server"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start server controller.",
	Run: func(cmd *cobra.Command, args []string) {
		serverService := server.NewServerService()
		err := serverService.Start()
		if err != nil {
			cmd.Println("Error starting server", err)
		}
	},
}
