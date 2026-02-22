package github

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseLink(link string) (*RepoInfo, error) {
	link = strings.TrimSuffix(link, "/")

	re := regexp.MustCompile(`^https://github\.com/([^/]+)/([^/]+)(?:/tree/([^/]+))?(?:/(.+))?$`)
	matches := re.FindStringSubmatch(link)
	if len(matches) != 5 {
		return nil, fmt.Errorf("invalid Link please try again")
	}
	return &RepoInfo{
		Owner:  matches[1],
		Repo:   matches[2],
		Branch: matches[3],
		Path:   matches[4],
	}, nil
}
