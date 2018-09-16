package cmdFlag

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin/json"
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

func PrintCommandFlag() {
	flag.StringVar(&name, "n", "xieWei", "show user name")
	flag.StringVar(&url, "u", "https://www.baidu.com", "show user url")
	flag.StringVar(&email, "e", "wuxiaoshen@shu.edu.cn", "show user email")
	flag.StringVar(&company, "c", "ReadSense", "show user company")
	flag.Parse()
	var oneUser userInfoFlag
	oneUser.Name = name
	oneUser.Url = url
	oneUser.Email = email
	oneUser.Company = company
	fmt.Println(name, url, email, company)
	jsonByte, _ := json.MarshalIndent(oneUser, " ", " ")
	fmt.Println(string(jsonByte))
}
