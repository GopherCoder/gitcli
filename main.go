package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "0.0.1"
	cliName = "gitcli"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Args:    cobra.ExactArgs(0),
	Aliases: []string{"-v"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("%s version: %s\n", cliName, version))
	},
}

var RootCmd = &cobra.Command{
	Use: cliName,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("%s is an application for operation github.\n", cliName))

	},
}

func Execute() {
	RootCmd.AddCommand(versionCmd)
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()

}
