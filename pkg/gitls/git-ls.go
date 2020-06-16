package gitls

import (
	"context"
	"net/url"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

const (
	firstPage  = 1
	lastPage   = 0
	maxPerPage = 100
)

type ghClient struct {
	gh    *github.Client
	Token string
}

// NewClient returns a GitHub client for the application to use
func NewClient(apiKey string) (client *ghClient) {
	client = &ghClient{
		Token: apiKey,
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apiKey},
	)
	tc := oauth2.NewClient(ctx, ts)
	client.gh = github.NewClient(tc)
	return
}

// NewEnterpriseClient returns a client with a custom github API base URL
func NewEnterpriseClient(baseURL, apiKey string) (client *ghClient) {
	client = NewClient(apiKey)
	base, _ := url.Parse(baseURL)
	client.gh.BaseURL = base
	return
}
