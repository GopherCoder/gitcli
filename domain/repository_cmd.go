package domain

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/gitcli/infrastructure"
	"github.com/wuxiaoxiaoshen/gitcli/infrastructure/errors"


	"github.com/tidwall/gjson"

	"github.com/spf13/cobra"
)

type RepoInfo struct {
	URL         string `json:"url"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Private     bool   `json:"private"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Watchers    string `json:"watchers"`
	Forks       string `json:"forks"`
	Language    string `json:"language"`
	SshUrl      string `json:"ssh_url"`
}

var RepoCmd = &cobra.Command{
	Use:     "repos",
	Aliases: []string{"r"},
	Short:   "show user repository",
	Long:    "show user all repositories or just one repository detail info",
	Run:     repoCommand,
}

func repoCommand(cmd *cobra.Command, args []string) {

	var url string

	url = makeURL(args)
	fmt.Println(url, args)

	if len(args) < 2 {
		fmt.Println("should add one more arguments")
		return
	}

	if url != "None" && args[1] == "all" {
		response, _ := infrastructure.GetResponseNetHttp(url)
		if ok := gjson.ParseBytes(response).IsArray(); ok != true {
			fmt.Println(&errors.ErrorCmdArray)
			return
		}

		var resultArray []RepoInfo
		for index, array := range gjson.ParseBytes(response).Array() {
			// default number 10
			if index >= 10 {
				break
			}
			var oneRepoInfo RepoInfo
			oneRepoInfo.URL = array.Get("html_url").String()
			oneRepoInfo.Name = array.Get("name").String()
			oneRepoInfo.FullName = array.Get("full_name").String()
			oneRepoInfo.Private = array.Get("private").Bool()
			oneRepoInfo.Description = array.Get("description").String()
			oneRepoInfo.CreatedAt = array.Get("created_at").String()
			oneRepoInfo.UpdatedAt = array.Get("updated_at").String()
			oneRepoInfo.Watchers = array.Get("watchers").String()
			oneRepoInfo.Forks = array.Get("forks").String()
			oneRepoInfo.Language = array.Get("language").String()
			oneRepoInfo.SshUrl = array.Get("ssh_url").String()
			resultArray = append(resultArray, oneRepoInfo)
		}

		jsonByte, _ := json.MarshalIndent(resultArray, " ", " ")
		fmt.Println(string(jsonByte))

	} else if args[1] != "all" {
		repoSingle(args)
	} else {
		fmt.Println("\ntry gitcli --help")
		return

	}

}

func repoSingle(args []string) {
	url := makeURL(args)
	response, _ := infrastructure.GetResponseNetHttp(url)

	var oneRepoInfo RepoInfo
	jsonRepoInfo := gjson.ParseBytes(response)
	oneRepoInfo.URL = jsonRepoInfo.Get("html_url").String()
	oneRepoInfo.Name = jsonRepoInfo.Get("name").String()
	oneRepoInfo.FullName = jsonRepoInfo.Get("full_name").String()
	oneRepoInfo.Private = jsonRepoInfo.Get("private").Bool()
	oneRepoInfo.Description = jsonRepoInfo.Get("description").String()
	oneRepoInfo.CreatedAt = jsonRepoInfo.Get("created_at").String()
	oneRepoInfo.UpdatedAt = jsonRepoInfo.Get("updated_at").String()
	oneRepoInfo.Watchers = jsonRepoInfo.Get("watchers").String()
	oneRepoInfo.Forks = jsonRepoInfo.Get("forks").String()
	oneRepoInfo.Language = jsonRepoInfo.Get("language").String()
	oneRepoInfo.SshUrl = jsonRepoInfo.Get("ssh_url").String()
	jsonByte, _ := json.MarshalIndent(oneRepoInfo, " ", " ")
	fmt.Println(string(jsonByte))

}

func makeURL(args []string) string {
	if len(args) == 2 && args[1] == "all" {
		return fmt.Sprintf(infrastructure.API["repo_url"], args[0])
	}

	if len(args) == 2 && args[1] != "all" {
		return fmt.Sprintf(infrastructure.API["repo_single_url"], args[0], args[1])
	}
	return "None"

}
