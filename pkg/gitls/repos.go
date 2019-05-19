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
	ctx := context.Background()
	opts := &github.RepositoryListOptions{}
	opts.PerPage = 100

	// TODO properly paginate when more that 100 records exist
	ghRepos, _, err := gls.gh.Repositories.List(ctx, "", opts)
	if err != nil {
		return
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
