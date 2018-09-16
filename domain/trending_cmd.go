package domain

import (
	"fmt"
	"gitcli/infrastructure"

	"github.com/spf13/cobra"
)

/*
1. 需要爬虫
2. 先爬取所有的语言
3. 再构造URL 爬取数据
4. 只返回Json 格式，表格太长，不适合展示
*/

type trendingRepo struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stars       string `json:"stars"`
	Forked      string `json:"forked"`
	Today       string `json:"today"`
}

var TrendingCmd = &cobra.Command{
	Use:   "trend",
	Short: "show github trending by language",
	Long:  "show github trending by language , period , and date, get more detail information",
	Run:   trendingCommand,
}

func trendingCommand(cmd *cobra.Command, args []string) {}

func showTrendingRepoByJson() {}

func showTrendingRepoByTable() {}

func makeTrendingUrl(args []string) string {
	var url string
	if len(args) < 1 {
		fmt.Println("you should at least add one argument")
		url = "None"
	} else if len(args) == 1 {
		url = fmt.Sprintf(infrastructure.API["trending_url"], args[0], "daily")
	} else {
		url = fmt.Sprintf(infrastructure.API["trending_url"], args[0], args[1])
	}
	return url
}

func getTrendingRepos(url string) []trendingRepo {
	response, _ := infrastructure.GetResponseNetHttp(url)
	fmt.Println(response)
	return nil

}
