package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wuxiaoxiaoshen/gitcli/infrastructure"
	"regexp"
	"strings"

	"github.com/alexeyco/simpletable"

	"github.com/PuerkitoBio/goquery"

	"github.com/spf13/cobra"
)

/*
1. 需要爬虫
2. 先爬取所有的语言
3. 再构造URL 爬取数据
4. 只返回Json 格式，表格太长，不太适合展示
5. 2019-12-09: github-trend 网页改版，重新编写
*/

type trendingRepo struct {
	Name        string `json:"name"`
	TotalStar   string `json:"total_star"`
	SinceStar   string `json:"since_star"`
	Fork        string `json:"fork"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Url         string `json:"url"`
}

var TrendingCmd = &cobra.Command{
	Use:   "trend",
	Short: "show github trending by language",
	Long:  "show github trending by language , period , and date, get more detail information",
	Run:   trendingCommand,
}

func trendingCommand(cmd *cobra.Command, args []string) {

	var url string
	url = makeTrendingUrl(args)
	if url == "None" {
		return
	}
	var trendingRepoInfo []trendingRepo
	trendingRepoInfo = getTrendingRepos(url)

	if args[len(args)-1] == "json" {
		showTrendingRepoByJson(trendingRepoInfo)
	}

	if args[len(args)-1] == "table" {
		showTrendingRepoByTable(trendingRepoInfo)
	}

}

func showTrendingRepoByJson(trendingRepos []trendingRepo) {
	jsonByte, _ := json.MarshalIndent(trendingRepos, " ", " ")
	fmt.Println(string(jsonByte))
}

func showTrendingRepoByTable(trendingRepos []trendingRepo) {
	table := simpletable.New()
	headers := []string{"name", "star", "now_star", "fork", "language", "url"}
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
	for _, item := range trendingRepos {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: item.Name},
			{Align: simpletable.AlignLeft, Text: item.TotalStar},
			{Align: simpletable.AlignLeft, Text: item.SinceStar},
			{Align: simpletable.AlignLeft, Text: item.Fork},
			{Align: simpletable.AlignLeft, Text: item.Language},
			{Align: simpletable.AlignLeft, Text: item.Url},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}
	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}

func makeTrendingUrl(args []string) string {
	var url string
	if len(args) < 1 {
		fmt.Println("you should at least add one argument")
		url = "None"
	} else if len(args) == 1 {
		url = fmt.Sprintf(infrastructure.API["trending_url"], strings.ToLower(args[0]), "daily")
	} else {
		url = fmt.Sprintf(infrastructure.API["trending_url"], strings.ToLower(args[0]), args[1])
	}
	return url
}

func getTrendingRepos(url string) []trendingRepo {
	var languagePattern string
	languagePattern = `trending/(.*?)\?`
	languageRegexp := regexp.MustCompile(languagePattern)
	allMatch := languageRegexp.FindAllStringSubmatch(url, -1)
	var trendingRepos []trendingRepo

	response, _ := infrastructure.GetResponseNetHttp(url)
	responseByte := bytes.NewReader(response)
	doc, _ := goquery.NewDocumentFromReader(responseByte)
	doc.Find("article.Box-row").Each(func(i int, selection *goquery.Selection) {
		name := strings.TrimSpace(selection.Find("h1 a").Text())
		newReplacer := strings.NewReplacer(" ", "", "\n", "", "\t", "")
		newName := newReplacer.Replace(name)

		description := strings.TrimSpace(selection.Find("p").Text())

		number := selection.Find("div").Eq(1)

		totalStar := strings.TrimSpace(number.Find("a").Eq(0).Text())

		fork := strings.TrimSpace(number.Find("a").Eq(1).Text())

		sinceStar := strings.TrimSpace(number.Find("span").Last().Text())

		pattern := "(.*?)stars"
		regexpPattern := regexp.MustCompile(pattern)
		all := regexpPattern.FindAllStringSubmatch(sinceStar, -1)
		var sinceNumber string
		for _, one := range all {
			sinceNumber = strings.TrimSpace(one[1])
		}

		var oneTrendingRepo trendingRepo
		oneTrendingRepo = trendingRepo{
			Name:        newName,
			Description: description,
			TotalStar:   totalStar,
			Fork:        fork,
			SinceStar:   sinceNumber,
			Language:    allMatch[0][1],
			Url:         "https://github.com/" + newName,
		}
		trendingRepos = append(trendingRepos, oneTrendingRepo)
	})
	return trendingRepos
}
