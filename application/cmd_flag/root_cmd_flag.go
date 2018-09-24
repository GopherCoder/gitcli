package cmdFlag

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

var (
	name    string
	url     string
	email   string
	company string
)

type userInfoFlag struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Email   string `json:"email"`
	Company string `json:"company"`
}

//func PrintCommandFlag() {
//	flag.StringVar(&name, "n", "xieWei", "show user name")
//	flag.StringVar(&url, "u", "https://www.baidu.com", "show user url")
//	flag.StringVar(&email, "e", "wuxiaoshen@shu.edu.cn", "show user email")
//	flag.StringVar(&company, "c", "ReadSense", "show user company")
//	//flag.Parse()
//	var oneUser userInfoFlag
//	oneUser.Name = name
//	oneUser.Url = url
//	oneUser.Email = email
//	oneUser.Company = company
//	jsonByte, _ := json.MarshalIndent(oneUser, " ", " ")
//	fmt.Println(string(jsonByte))
//}

func WeatherStorager(city string) ([]byte, error) {
	//api := "https://www.sojson.com/open/api/weathera/json.shtml?city=%s"
	api := "http://wthrcdn.etouch.cn/weather_mini?city=%s"
	url := fmt.Sprintf(api, city)
	//url := "http://www.weather.com.cn/data/sk/101010100.html"
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	fmt.Println(response.StatusCode)
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	//var results interface{}
	//json.Unmarshal(result, &results)
	fmt.Println(string([]byte(result)))
	return ioutil.ReadAll(response.Body)

}

type GithubAccountInfo struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowerURL       string `json:"follower_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	Name              string `json:"name"`
	Company           string `json:"company"`
	Blog              string `json:"blog"`
	Location          string `json:"location"`
	Email             string `json:"email"`
	Hireable          bool   `json:"hireable"`
	Bio               string `json:"bio"`
	PublicRepos       int    `json:"public_repos"`
	PublicGists       int    `json:"public_gists"`
	Followers         int    `json:"followers"`
	Following         int    `json:"following"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

func GithubUserFields() {
	var fields []string
	account := GithubAccountInfo{}
	s := reflect.TypeOf(&account).Elem()
	for i := 0; i < s.NumField(); i++ {
		fields = append(fields, s.Field(i).Name)
	}
	fieldsJson, _ := json.MarshalIndent(fields, " ", "")
	fmt.Println(string(fieldsJson))
}

func GithubUserStorager(name string) GithubAccountInfo {
	url := fmt.Sprintf("https://api.github.com/users/%s", name)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return GithubAccountInfo{}
	}
	fmt.Println(response.StatusCode)
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	fmt.Println(string([]byte(result)))
	var account GithubAccountInfo
	err = json.Unmarshal([]byte(result), &account)
	if err != nil {
		fmt.Println(err)
		return GithubAccountInfo{}
	}
	return account

}
