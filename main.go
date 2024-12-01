package main

import (
	"fmt"
	"github.com/AmadlaOrg/hery-playground/server/cmd"
	"github.com/spf13/cobra"
	"log"
)

const appName = "hery-playground"
const appTitleName = "HERY - Playground"
const version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:     appName,
	Short:   appTitleName + " CLI application",
	Version: version,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + appName,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(appName + " version " + version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(cmd.StartCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		//os.Exit(1)
	}
}
