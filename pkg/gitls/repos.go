package gitls

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/go-github/v25/github"
)

type repo struct {
	Name            string
	Descripiton     string
	URL             string
	Owner           string
	Organization    string
	StargazersCount int
	Private         bool
}

func (gls *ghClient) Repos() {
	allRepos, err := gls.repos()
	if err != nil {
		log.Fatal(err)
	}

	out, err := json.MarshalIndent(allRepos, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

func (gls *ghClient) repos() (all []*repo, err error) {
	var resp *github.Response
	var ghRepos []*github.Repository

	ctx := context.Background()
	opts := &github.RepositoryListOptions{}
	opts.PerPage = maxPerPage

	ghRepos, resp, err = gls.gh.Repositories.List(ctx, "", opts)
	if err != nil {
		return
	}

	opts.Page = resp.NextPage
	for resp.NextPage != 0 {
		rps, resp, err := gls.gh.Repositories.List(ctx, "", opts)
		if err != nil {
			return all, err
		}
		opts.Page = resp.NextPage
		ghRepos = append(ghRepos, rps...)
		if resp.NextPage == 0 {
			break
		}
	}

	for _, r := range ghRepos {
		re := &repo{
			Name:            r.GetName(),
			Descripiton:     r.GetDescription(),
			URL:             r.GetURL(),
			Owner:           r.GetOwner().GetLogin(),
			Organization:    r.GetOrganization().GetName(),
			StargazersCount: r.GetStargazersCount(),
			Private:         r.GetPrivate(),
		}
		all = append(all, re)
	}
	return
}
