package gitls

import (
	"log"

	"gopkg.in/src-d/go-git.v4"
)

// Plunder will find all private repositories from a given token and clone them
// into a given folder.
func (gls *ghClient) Plunder() {
	allRepos, err := gls.repos()
	if err != nil {
		log.Fatal(err)
	}

	var privateRepos []*repo
	for _, repo := range allRepos {
		if repo.Private {
			privateRepos = append(privateRepos, repo)
		}
	}

	for _, repo := range privateRepos {
		cloneURL := repo.CloneableURL(gls.Token)
		name := repo.Name

		log.Println("Cloning " + name)
		_, _, err := CloneRepository(cloneURL, name)
		if err != nil {
			log.Println(err)
		}
	}
}

// CloneRepository clones a repository
func CloneRepository(url, dest string) (*git.Repository, string, error) {
	repository, err := git.PlainClone(dest, false, &git.CloneOptions{URL: url})
	if err != nil {
		return nil, dest, err
	}

	return repository, dest, err
}
