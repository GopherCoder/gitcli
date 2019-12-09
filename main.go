package main

import (
	"flag"
	"github.com/wuxiaoxiaoshen/gitcli/application/cmd_cobra"
	cmdFlag "github.com/wuxiaoxiaoshen/gitcli/application/cmd_flag"
	cmdOs "github.com/wuxiaoxiaoshen/gitcli/application/cmd_os"
)

func OSHelper() {
	cmdOs.PrintCmdOs()
}
func FlagHelper() {
	var Account string
	flag.StringVar(&Account, "a", "wuxiaoxiaoshen", "show github account user info fields")
	flag.Parse()
	// 如果 Account 输入是 field 返回结构体所有字段，否则返回用户信息
	if Account == "field" {
		cmdFlag.GithubUserFields()
		return
	} else {
		cmdFlag.GithubUserStorager(Account)
		return
	}
}
func CobraHelper() {
	cmdCobra.Execute()
}

func main() {
	//OSHelper()
	//FlagHelper()
	CobraHelper()
}
