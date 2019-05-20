package gitls

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/go-github/v25/github"
)

type gist struct {
	Owner       string
	Descripiton string
	GitPullURL  string
	Files       []github.GistFilename
	Private     bool
}

func (gls *ghClient) Gists() {
	allGists, err := gls.gists()
	if err != nil {
		log.Fatal(err)
	}

	out, err := json.MarshalIndent(allGists, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

func (gls *ghClient) gists() (all []*gist, err error) {
	var resp *github.Response
	var ghGists []*github.Gist

	ctx := context.Background()
	opts := &github.GistListOptions{}
	opts.PerPage = maxPerPage

	ghGists, resp, err = gls.gh.Gists.List(ctx, "", opts)
	if err != nil {
		return
	}

	opts.Page = resp.NextPage
	for resp.NextPage != 0 {
		rps, resp, err := gls.gh.Gists.List(ctx, "", opts)
		if err != nil {
			return all, err
		}
		opts.Page = resp.NextPage
		ghGists = append(ghGists, rps...)
		if resp.NextPage == 0 {
			break
		}
	}

	for _, g := range ghGists {
		var filenames []github.GistFilename
		for name, _ := range g.Files {
			filenames = append(filenames, name)
		}
		gst := &gist{
			Owner:       g.GetOwner().GetLogin(),
			Descripiton: g.GetDescription(),
			GitPullURL:  g.GetGitPullURL(),
			Private:     !g.GetPublic(),
			Files:       filenames,
		}
		all = append(all, gst)
	}
	return
}
