package domain

import "github.com/spf13/cobra"

var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "show user information",
	Long: `information contains username、url、
			html_url、name、company、location、
			public_repos、followers、following、created_at、updated_at`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func showUserInfoJson(cmd *cobra.Command, args []string) {

}

func showUserInfoTable() {}
