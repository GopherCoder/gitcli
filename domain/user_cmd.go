package domain

import (
	"fmt"
	"gitcli/infrastructure"

	"github.com/tidwall/gjson"

	"github.com/spf13/cobra"
)

// UserCmd ...
var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "show user information",
	Long: `information contains username、url、
			html_url、name、company、location、
			public_repos、followers、following、created_at、updated_at`,
	Args: cobra.MinimumNArgs(1),
	Run:  UserCommand,
}

// UserCommand ...
func UserCommand(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println(fmt.Sprintf("try %s --help", cmd.Use))
		return
	}

	url := fmt.Sprintf(infrastructure.API["user_url"], args[0])
	fmt.Println(url)
	response, _ := infrastructure.GetResponseNetHttp(url)
	responseResult := gjson.ParseBytes(response)

	switch args[1] {

	case "name":
		fmt.Println(responseResult.Get("name"))
	case "url":
		fmt.Println(responseResult.Get("html_url"))
	case "location":
		fmt.Println(responseResult.Get("location"))
	case "create":
		fmt.Println(responseResult.Get("created_at"))
	case "update":
		fmt.Println(responseResult.Get("updated_at"))
	case "followers":
		fmt.Println(responseResult.Get("followers"))
	case "following":
		fmt.Println(responseResult.Get("following"))
	case "repos":
		fmt.Println(responseResult.Get("public_repos"))
	case "bio":
		fmt.Println(responseResult.Get("bio"))
	case "email":
		fmt.Println(responseResult.Get("email"))
	case "company":
		fmt.Println(responseResult.Get("company"))
	case "all":
		fmt.Println(responseResult.String())
	}

}

func showUserInfoJson(arg string) {

}

func showUserInfoTable(arg string) {}

func showUserName(result gjson.Result) {

}
