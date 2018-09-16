package cmdOs

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin/json"
)

type userInfoOs struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company"`
}

func PrintCmdOs() {
	args := os.Args
	if len(args) != 4 {
		fmt.Println("you need add name,email,company field")
		return
	}
	var oneUserInfoOs userInfoOs
	oneUserInfoOs.Name = os.Args[1]
	oneUserInfoOs.Email = os.Args[2]
	oneUserInfoOs.Company = os.Args[3]

	jsonByte, _ := json.MarshalIndent(oneUserInfoOs, " ", " ")
	fmt.Println(string(jsonByte))
}
