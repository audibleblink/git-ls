package gitls

import (
	"context"

	"github.com/google/go-github/v25/github"
	"golang.org/x/oauth2"
)

const (
	firstPage  = 1
	lastPage   = 0
	maxPerPage = 100
)

type ghClient struct {
	gh *github.Client
}

// NewClient returns a GitHub client for the application to use
func NewClient(apiKey string) (client *ghClient) {
	client = &ghClient{}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apiKey},
	)
	tc := oauth2.NewClient(ctx, ts)
	client.gh = github.NewClient(tc)
	return
}
