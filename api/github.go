// Package githuber gives a cli interface for github v3 api
package githuber

import "time"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
