package domain

import (
	"fmt"
	"gitcli/infrastructure"
	"gitcli/infrastructure/errors"
	"strings"

	"github.com/alexeyco/simpletable"

	"github.com/gin-gonic/gin/json"

	"github.com/tidwall/gjson"

	"github.com/spf13/cobra"
)

type Follower struct {
	Name           string `json:"name"`
	Url            string `json:"url"`
	FollowerNumber string `json:"follower_number"`
	Following      string `json:"following"`
	Company        string `json:"company"`
	Email          string `json:"email"`
}

var FollowersCmd = &cobra.Command{
	Use:     "follower",
	Short:   "show user follower info",
	Long:    "show user follower detail info, such as url, name , full_name, repos, and so on",
	Aliases: []string{"f"},
	Run:     followerCommand,
}

func followerCommand(cmd *cobra.Command, args []string) {
	fmt.Println(cmd.Use, args)
	url := makeUserFollowerURL(args)
	fmt.Println(url)
	response, _ := infrastructure.GetResponseNetHttp(makeUserFollowerURL(args))
	if ok := gjson.ParseBytes(response).IsArray(); ok != true {
		fmt.Println(&errors.ErrorCmdArray)
		return
	}
	if args[1] == "json" {
		showFollowersJson(gjson.ParseBytes(response))
	}
	if args[1] == "table" {
		showFollowerTable(gjson.ParseBytes(response))
	}

}

func makeUserFollowerURL(args []string) string {

	if len(args) >= 1 {
		return fmt.Sprintf(infrastructure.API["user_follower_url"], args[0])
	}
	return "None"
}

func getFollowerResult(result gjson.Result) []Follower {
	var arrayFollowers []Follower

	for _, array := range result.Array() {
		var oneFollower Follower
		oneFollower.Name = array.Get("login").String()
		oneFollower.Url = array.Get("html_url").String()

		url := array.Get("url").String()
		partInfo := getPartInfo(url)

		oneFollower.Following = partInfo["following"]
		oneFollower.Company = partInfo["company"]
		oneFollower.Email = partInfo["email"]
		oneFollower.FollowerNumber = partInfo["followers"]

		arrayFollowers = append(arrayFollowers, oneFollower)

	}
	return arrayFollowers
}

func showFollowersJson(result gjson.Result) {
	arrayFollowers := getFollowerResult(result)
	jsonByte, _ := json.MarshalIndent(arrayFollowers, " ", " ")
	fmt.Println(string(jsonByte))

}

func showFollowerTable(result gjson.Result) {

	headers := []string{"name", "url", "followers", "following", "email", "company"}

	infos := getFollowerResult(result)

	table := simpletable.New()

	var cells []*simpletable.Cell

	for _, header := range headers {
		cell := &simpletable.Cell{
			Align: simpletable.AlignCenter, Text: strings.ToUpper(header),
		}
		cells = append(cells, cell)
	}
	table.Header = &simpletable.Header{
		Cells: cells,
	}
	for _, info := range infos {
		cell := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: info.Name},
			{Align: 0, Text: info.Url},
			{Align: 0, Text: info.FollowerNumber},
			{Align: 0, Text: info.Following},
			{Align: 0, Text: info.Email},
			{Align: 0, Text: info.Company},
		}
		table.Body.Cells = append(table.Body.Cells, cell)
	}
	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())

}

func getPartInfo(url string) map[string]string {

	var mapValue = make(map[string]string)
	response, _ := infrastructure.GetResponseNetHttp(url)
	mapValue["following"] = gjson.ParseBytes(response).Get("following").String()
	mapValue["company"] = gjson.ParseBytes(response).Get("company").String()
	mapValue["email"] = gjson.ParseBytes(response).Get("email").String()
	mapValue["followers"] = gjson.ParseBytes(response).Get("followers").String()
	return mapValue

}
