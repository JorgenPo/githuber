package githuber

import (
	"net/http"
	"fmt"
	"bytes"
	"os"
	"encoding/json"
	"time"
)

const authUrl = "https://api.github.com/authorized"

type Application struct {
	Url			string
	Name		string
	ClientId	string
}

type Authorization struct {
	Id			int
	Url 		string
	Scopes 		[]string
	Token 		string
	HashedToken	string
	App 		*Application
	Note		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

// Authorize function creates a user authorization token with
// OAuth2 github api
func Authorize(clientId, clientSecret string) (Authorization, error) {
	var content =
		` { note: "user script", client_id: %s, client_secret: %s }`

	content = fmt.Sprintf(content, clientId, clientSecret)

	resp, err := http.Post(authUrl, "application/json", bytes.NewReader([]byte(content)))
	if err != nil {
		// TODO: move error messages to invoker code
		fmt.Fprintf(os.Stderr, "githuber authorization error: failed to make http post query: %v", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		// TODO: move error messages to invoker code
		fmt.Fprintf(os.Stderr, "githuber authorization error: server response with status %s (%d)", resp.Status, resp.StatusCode)
		return nil, err
	}

	var auth Authorization
	if err := json.NewDecoder(resp.Body).Decode(&auth); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()

	return auth, nil
}