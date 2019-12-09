package cmdCobra

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/gitcli/domain"
	"github.com/wuxiaoxiaoshen/gitcli/infrastructure"
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

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "show network status",
	Long:  "connect to internet because of cli should link to internet",
	Run:   connectInternet,
}

func connectInternet(cmd *cobra.Command, args []string) {
	if len(args) > 1 {
		fmt.Println(fmt.Sprintf("try %s %s--help"), cliName, cmd.Use)
		return
	}
	url := "https://www.baidu.com"
	code := infrastructure.InternetStatus(url)
	if code == 200 {
		fmt.Println("Connect internet Success")
	} else {
		fmt.Println("Connect internet Failed")
	}

}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "make basic auth token",
	Long:  "you need init command and get basic auth token",
	Run:   getBasicAuthToken,
}

func getBasicAuthToken(cmd *cobra.Command, args []string) {
	//token := args[0] + ":" + args[1]

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

	// version command
	RootCmd.AddCommand(versionCmd)
	// status command
	RootCmd.AddCommand(statusCmd)
	// user command
	RootCmd.AddCommand(domain.UserCmd)
	// repos command
	RootCmd.AddCommand(domain.RepoCmd)
	// followers command
	RootCmd.AddCommand(domain.FollowersCmd)
	// search command
	RootCmd.AddCommand(domain.SearchCmd)
	// trending command
	RootCmd.AddCommand(domain.TrendingCmd)
	// organization command
	RootCmd.AddCommand(domain.OrganizationCmd)
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
