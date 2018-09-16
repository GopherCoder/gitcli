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

type organizationRepo struct {
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	URL           string `json:"html_url"`
	Description   string `json:"description"`
	WatchersCount string `json:"watchers_count"`
	Language      string `json:"language"`
	ForksCount    string `json:"forks_count"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

var OrganizationCmd = &cobra.Command{
	Use:   "organization",
	Short: "show organization info",
	Long:  "show organization detail info such as repo (fork, description, watch), ,member, and so on",
	Run:   organizationCommand,
}

func organizationCommand(cmd *cobra.Command, args []string) {

	url := makeOrganizationURL(args)
	organizationRepos := getOrganizationRepo(url)

	if args[len(args)-1] == "json" {
		showOrganizationByJson(organizationRepos)
	}
	if args[len(args)-1] == "table" {
		showOrganizationByTable(organizationRepos)
	}
}

func showOrganizationByJson(organizations []organizationRepo) {
	jsonByte, _ := json.MarshalIndent(organizations, "", "\t")
	fmt.Println(string(jsonByte))
}

func showOrganizationByTable(organizations []organizationRepo) {
	table := simpletable.New()

	headers := []string{"name", "full_name", "url", "description", "watch_count", "language", "fork_count",
		"created_at", "updated_at"}
	var cells []*simpletable.Cell

	for _, header := range headers {
		cell := &simpletable.Cell{
			Align: simpletable.AlignLeft, Text: strings.ToUpper(header),
		}
		cells = append(cells, cell)
	}
	table.Header = &simpletable.Header{
		Cells: cells,
	}

	for _, item := range organizations {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: item.Name},
			{Align: simpletable.AlignLeft, Text: item.FullName},
			{Align: simpletable.AlignLeft, Text: item.URL},
			{Align: simpletable.AlignLeft, Text: item.Description},
			{Align: simpletable.AlignLeft, Text: item.WatchersCount},
			{Align: simpletable.AlignLeft, Text: item.Language},
			{Align: simpletable.AlignLeft, Text: item.ForksCount},
			{Align: simpletable.AlignLeft, Text: item.CreatedAt},
			{Align: simpletable.AlignLeft, Text: item.UpdatedAt},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}
	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}

func makeOrganizationURL(args []string) string {
	var url string
	if len(args) < 1 {
		fmt.Println("you should add at least one argument")
		url = "None"
	} else {
		url = fmt.Sprintf(infrastructure.API["organization_repo_url"], args[0])
	}
	return url
}

func getOrganizationRepo(url string) []organizationRepo {
	response, _ := infrastructure.GetResponseNetHttp(url)

	if ok := gjson.ParseBytes(response).IsArray(); ok != true {
		fmt.Println(&errors.ErrorCmdArray)
		return nil
	}
	var organizationRepos []organizationRepo

	for _, array := range gjson.ParseBytes(response).Array() {
		var oneOrganizationRepo organizationRepo
		oneOrganizationRepo = organizationRepo{
			Name:          array.Get("name").String(),
			FullName:      array.Get("full_name").String(),
			URL:           array.Get("html_url").String(),
			Description:   array.Get("description").String(),
			WatchersCount: array.Get("watchers_count").String(),
			ForksCount:    array.Get("forks_count").String(),
			Language:      array.Get("language").String(),
			CreatedAt:     array.Get("created_at").String(),
			UpdatedAt:     array.Get("updated_at").String(),
		}
		organizationRepos = append(organizationRepos, oneOrganizationRepo)
	}
	return organizationRepos

}
