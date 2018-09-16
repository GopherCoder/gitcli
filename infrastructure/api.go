package infrastructure

var API = map[string]string{
	"current_user_url":                     "https://api.github.com/user",
	"current_user_authorizations_html_url": "https://github.com/settings/connections/applications{/client_id}",
	"authorizations_url":                   "https://api.github.com/authorizations",
	"code_search_url":                      "https://api.github.com/search/code?q={query}{&page,per_page,sort,order}",
	"commit_search_url":                    "https://api.github.com/search/commits?q={query}{&page,per_page,sort,order}",
	"emails_url":                           "https://api.github.com/user/emails",
	"emojis_url":                           "https://api.github.com/emojis",
	"events_url":                           "https://api.github.com/events",
	"feeds_url":                            "https://api.github.com/feeds",
	"followers_url":                        "https://api.github.com/user/followers",
	"following_url":                        "https://api.github.com/user/following{/target}",
	"gists_url":                            "https://api.github.com/gists{/gist_id}",
	"hub_url":                              "https://api.github.com/hub",
	"issue_search_url":                     "https://api.github.com/search/issues?q={query}{&page,per_page,sort,order}",
	"issues_url":                           "https://api.github.com/issues",
	"keys_url":                             "https://api.github.com/user/keys",
	"notifications_url":                    "https://api.github.com/notifications",
	"organization_repositories_url":        "https://api.github.com/orgs/{org}/repos{?type,page,per_page,sort}",
	"organization_url":                     "https://api.github.com/orgs/{org}",
	"public_gists_url":                     "https://api.github.com/gists/public",
	"rate_limit_url":                       "https://api.github.com/rate_limit",
	"repository_url":                       "https://api.github.com/repos/{owner}/{repo}",
	"current_user_repositories_url":        "https://api.github.com/user/repos{?type,page,per_page,sort}",
	"starred_url":                          "https://api.github.com/user/starred{/owner}{/repo}",
	"starred_gists_url":                    "https://api.github.com/gists/starred",
	"team_url":                             "https://api.github.com/teams",
	"user_organizations_url":               "https://api.github.com/user/orgs",
	"user_repositories_url":                "https://api.github.com/users/{user}/repos{?type,page,per_page,sort}",
	"user_search_url":                      "https://api.github.com/search/users?q={query}{&page,per_page,sort,order}",

	// userCommand
	"user_url": "https://api.github.com/users/%s",

	// repoCommand
	"repo_url":        "https://api.github.com/users/%s/repos",
	"repo_single_url": "https://api.github.com/repos/%s/%s",

	// followerCommand
	"user_follower_url": "https://api.github.com/users/%s/followers",

	// searchCommand
	"repository_search_url": "https://api.github.com/search/repositories?q=%s&page=%d&per_page=%d",
}
