package cmdCobra

import (
	"fmt"
	"gitcli/domain"
	"gitcli/infrastructure"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = infrastructure.Version
	cliName = infrastructure.ApplicationCmdName
)

// versionCmd ...
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "show application version",
	Args:    cobra.ExactArgs(0),
	Aliases: []string{"v"},
	Run:     printVersion,
}

// printVersion ...
func printVersion(cmd *cobra.Command, args []string) {
	if len(args) > 1 {
		fmt.Println(fmt.Sprintf("try %s %s --help", cliName, cmd.Use))
		return
	}
	fmt.Println(fmt.Sprintf("%s version: %s\n", cliName, version))
}

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:  cliName,
	Args: cobra.ExactArgs(0),
	Run:  printCliName,
}

// printCliName
func printCliName(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		fmt.Println(fmt.Sprintf("try %s --help", cmd.Use))
		return
	}
	fmt.Println(fmt.Sprintf("%s is an application for operation github.\n", cliName))
}

// Execute ...
func Execute() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(domain.UserCmd)
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
