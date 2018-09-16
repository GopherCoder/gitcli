package domain

import (
	"github.com/spf13/cobra"
)

var TrendingCmd = &cobra.Command{
	Use:   "trend",
	Short: "show github trending by language",
	Long:  "show github trending by language , period , and date, get more detail information",
	Run:   trendingCommand,
}

func trendingCommand(cmd *cobra.Command, args []string) {}

func showTrendingRepoByJson() {}

func showTrendingRepoByTable() {}
