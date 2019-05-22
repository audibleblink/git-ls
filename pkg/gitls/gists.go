package gitls

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

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
	var ghGists []*github.Gist
	ghGists = gls.gistPager(ghGists, firstPage)

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

func (gls *ghClient) gistPager(data []*github.Gist, page int) []*github.Gist {
	if page == lastPage {
		return data
	}

	ctx := context.Background()
	opts := &github.GistListOptions{}
	opts.PerPage = maxPerPage
	opts.Page = page

	gists, response, err := gls.gh.Gists.List(ctx, "", opts)
	if err != nil {
		return data
	}
	data = append(data, gists...)
	return gls.gistPager(data, response.NextPage)
}
