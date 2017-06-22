package githuber

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const issuesUrl = "https://api.github.com/search/issues"

// SearchIssues запрашивает github
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))

	resp, err := http.Get(issuesUrl + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Http get request error: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
