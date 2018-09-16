package domain

import (
	"encoding/json"
	"fmt"
	"gitcli/infrastructure"

	"github.com/alexeyco/simpletable"

	"github.com/tidwall/gjson"

	"github.com/spf13/cobra"
)

type Info struct {
	Field  string `json:"field"`
	Result string `json:"result"`
}

// UserCmd ...
var UserCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "show user information",
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

	if args[0] == "fields" {
		var fields = []string{"url", "login", "location", "created_at", "updated_at", "followers", "following", "public_repos",
			"bio", "email", "company"}
		jsonByte, _ := json.MarshalIndent(fields, " ", " ")
		fmt.Println(string(jsonByte))
		return
	}

	url := fmt.Sprintf(infrastructure.API["user_url"], args[0])

	response, _ := infrastructure.GetResponseNetHttp(url)
	responseResult := gjson.ParseBytes(response)

	if args[1] == "all" {
		jsonByte, _ := json.MarshalIndent(responseResult.Raw, " ", " ")
		fmt.Println(string(jsonByte))
		return
	}
	//fmt.Println(args)
	if len(args) <= 2 {
		fmt.Println("Should add one more argument")
		return
	}

	if args[1] != "all" && args[2] == "json" {

		infoField := showUserField(args[1], responseResult.Get(args[1]))
		if infoField.Result == "" {
			infoField.Result = "None"
		}
		showUserInfoJson(infoField)
		return
	}
	if args[1] != "all" && args[2] == "table" {
		infoField := showUserField(args[1], responseResult.Get(args[1]))
		showUserInfoTable(infoField)
		return
	}

}

// showUserInfoJson ...
func showUserInfoJson(info *Info) {
	jsonByte, _ := json.MarshalIndent(info, " ", " ")
	fmt.Println(string(jsonByte))

}

// showUserInfoTable ...
func showUserInfoTable(info *Info) {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "FIELD"},
			{Align: simpletable.AlignCenter, Text: "VALUE"},
		},
	}
	r := []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%s", info.Field)},
		{Align: simpletable.AlignLeft, Text: info.Result},
	}
	table.Body.Cells = append(table.Body.Cells, r)
	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())

}

// showUserField ...
func showUserField(field string, result gjson.Result) *Info {
	return &Info{
		Field:  field,
		Result: result.Raw,
	}

}
