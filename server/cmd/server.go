package cmd

import "github.com/spf13/cobra"

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start server controller.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
