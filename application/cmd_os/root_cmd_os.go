package cmdOs

import (
	"fmt"
	"html/template"
	"os"
)

type userInfoOs struct {
	File    string `json:"file"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company"`
}

func (u *userInfoOs) Template() {
	t := template.New("New Template for book")
	t, _ = t.Parse(`
An example of os cli.

Show User Information by template:
	FileName: {{.File}}
	Name: {{.Name}}
	Email: {{.Email}}
	Company: {{.Company}}

Use "user help <topic>" for more information about that topic.
`)
	t.Execute(os.Stdout, u)
}

func PrintCmdOs() {
	args := os.Args
	if len(args) != 4 {
		fmt.Println("you need add name,email,company field")
		return
	}
	var oneUserInfoOs userInfoOs
	oneUserInfoOs.File = os.Args[0]
	oneUserInfoOs.Name = os.Args[1]
	oneUserInfoOs.Email = os.Args[2]
	oneUserInfoOs.Company = os.Args[3]

	// json
	//jsonByte, _ := json.MarshalIndent(oneUserInfoOs, " ", " ")
	//fmt.Println(string(jsonByte))

	// template
	oneUserInfoOs.Template()
}
