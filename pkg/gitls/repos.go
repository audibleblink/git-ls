package gitls

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"

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

func (r *repo) CloneableURL(token string) (cloneURL string) {
	re1, _ := regexp.Compile("api\\.")
	re2, _ := regexp.Compile("repos/")

	temp := re1.ReplaceAllLiteralString(r.URL, fmt.Sprintf("%s@", token))
	cloneURL = re2.ReplaceAllLiteralString(temp, "")
	return
}

func (gls *ghClient) Collabs() {
	allRepos, err := gls.repos()
	if err != nil {
		log.Fatal(err)
	}

	tokenOwner := gls.TokenOwner()
	var filteredRepos []*repo
	for _, r := range allRepos {
		if r.Owner != tokenOwner {
			filteredRepos = append(filteredRepos, r)
		}
	}

	out, err := json.MarshalIndent(filteredRepos, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
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
	var ghRepos []*github.Repository
	ghRepos = gls.repoPager(ghRepos, firstPage)

	for _, ghRepo := range ghRepos {
		repository := &repo{
			Name:            ghRepo.GetName(),
			Descripiton:     ghRepo.GetDescription(),
			URL:             ghRepo.GetURL(),
			Owner:           ghRepo.GetOwner().GetLogin(),
			Organization:    ghRepo.GetOrganization().GetName(),
			StargazersCount: ghRepo.GetStargazersCount(),
			Private:         ghRepo.GetPrivate(),
		}
		all = append(all, repository)
	}
	return
}

func (gls *ghClient) repoPager(data []*github.Repository, page int) []*github.Repository {
	if page == lastPage {
		return data
	}

	ctx := context.Background()
	opts := &github.RepositoryListOptions{}
	opts.PerPage = maxPerPage
	opts.Page = page

	repos, response, err := gls.gh.Repositories.List(ctx, "", opts)
	if err != nil {
		return data
	}
	data = append(data, repos...)
	return gls.repoPager(data, response.NextPage)
}
