package domain

import (
	"fmt"
	"gitcli/infrastructure"
	"gitcli/infrastructure/errors"

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

func organizationCommand(cmd *cobra.Command, args []string) {}

func showOrganizationByJson() {}

func showOrganizationByTable() {}

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
