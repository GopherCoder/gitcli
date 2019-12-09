package domain

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/gitcli/infrastructure"
	"strconv"
	"strings"

	"github.com/alexeyco/simpletable"


	"github.com/tidwall/gjson"

	"github.com/spf13/cobra"
)

type item struct {
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	Url             string `json:"html_url"`
	Description     string `json:"description"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	Language        string `json:"language"`
	WatchersCount   string `json:"watchers_count"`
	ForksCount      string `json:"forks_count"`
	OpenIssuesCount string `json:"open_issues_count"`
	License         string `json:"license"`
}

var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "search repository from github",
	Long:  "search repository from github by keyword, query arguments contain q, page and per_page",
	Run:   searchRepository,
}

func searchRepository(cmd *cobra.Command, args []string) {

	var (
		url   string
		items []item
	)
	url = makeSearchRepoUrl(args)
	items = getSearchRepoResult(url)
	// json
	if args[len(args)-1] == "json" {
		showSearchRepoByJson(items)
	}
	// table
	if args[len(args)-1] == "table" {
		showSearchRepoByTable(items)
	}

}

func makeSearchRepoUrl(args []string) string {

	var url string
	if len(args) <= 1 {
		fmt.Println("you should add at least one argument")
		url = "None"
	} else if len(args) == 2 {
		url = fmt.Sprintf(infrastructure.API["repository_search_url"], args[0], 1, 10)
	} else if len(args) == 3 {
		url = fmt.Sprintf(infrastructure.API["repository_search_url"], args[0], args[1], 10)
	} else {
		page, _ := strconv.Atoi(args[1])
		perPage, _ := strconv.Atoi(args[2])
		url = fmt.Sprintf(infrastructure.API["repository_search_url"], args[0], page, perPage)
	}
	return url
}

func getSearchRepoResult(url string) []item {
	response, _ := infrastructure.GetResponseNetHttp(url)
	var items []item
	for _, array := range gjson.ParseBytes(response).Get("items").Array() {
		var oneItem item
		oneItem = item{
			Name:            array.Get("name").String(),
			FullName:        array.Get("full_name").String(),
			Url:             array.Get("html_url").String(),
			Description:     array.Get("description").String(),
			CreatedAt:       array.Get("created_at").String(),
			UpdatedAt:       array.Get("updated_at").String(),
			Language:        array.Get("language").String(),
			WatchersCount:   array.Get("watchers_count").String(),
			ForksCount:      array.Get("forks_count").String(),
			OpenIssuesCount: array.Get("open_issues_count").String(),
			License:         array.Get("license").String(),
		}
		items = append(items, oneItem)

	}
	return items

}

func showSearchRepoByJson(items []item) {
	jsonByte, _ := json.MarshalIndent(items, " ", " ")
	fmt.Println(string(jsonByte))
}

func showSearchRepoByTable(items []item) {
	table := simpletable.New()
	headers := []string{"name", "full_name", "html_url", "description", "created_at",
		"updated_at", "language", "watchers_count", "forks_count", "open_issues_count", "license"}

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

	for _, item := range items {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: item.Name},
			{Align: simpletable.AlignLeft, Text: item.FullName},
			{Align: simpletable.AlignLeft, Text: item.Url},
			{Align: simpletable.AlignLeft, Text: item.Description},
			{Align: simpletable.AlignLeft, Text: item.CreatedAt},
			{Align: simpletable.AlignLeft, Text: item.UpdatedAt},
			{Align: simpletable.AlignLeft, Text: item.Language},
			{Align: simpletable.AlignLeft, Text: item.WatchersCount},
			{Align: simpletable.AlignLeft, Text: item.ForksCount},
			{Align: simpletable.AlignLeft, Text: item.OpenIssuesCount},
			{Align: simpletable.AlignLeft, Text: item.License},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}
	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}
