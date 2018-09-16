package domain

import "github.com/spf13/cobra"

var Repo = &cobra.Command{
	Use:     "repos",
	Aliases: []string{"r"},
	Short:   "show user repository",
	Long:    "show user all repositories or just one repository detail info",
	Run:     repoCommand,
}

func repoCommand(cmd *cobra.Command, args []string) {}
