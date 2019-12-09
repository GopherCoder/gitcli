package domain

import (
	"fmt"
	"testing"
)

func TestTrend(t *testing.T) {
	fmt.Println(getTrendingRepos("https://github.com/trending/go?since=daily"))
}
