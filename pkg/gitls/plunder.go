package gitls

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"gopkg.in/src-d/go-git.v4"
)

// Plunder will find all private repositories from a given token and clone them
// into a given folder.
func (gls *ghClient) Plunder(privateOnly bool) {
	owner := gls.TokenOwner()
	charSet := spinner.CharSets[14]
	spnTime := 100 * time.Millisecond

	msg := fmt.Sprintf("Seaching for private repos to which %s has access", owner)
	final := fmt.Sprintf("%s\nFound the following repos:\n", msg)

	spin := spinner.New(charSet, spnTime, setOpts(msg, final))
	spin.Start()

	allRepos, err := gls.repos()
	if err != nil {
		log.Fatal(err)
	}
	spin.Stop()

	var reposToClone []*repo
	if privateOnly {
		for _, repo := range allRepos {
			if repo.Private {
				fmt.Printf("%s/%s\n", repo.Owner, repo.Name)
				reposToClone = append(reposToClone, repo)
			}
		}
	} else {
		reposToClone = allRepos[:]
	}

	var wg sync.WaitGroup
	wg.Add(len(reposToClone))

	final = stats(owner, reposToClone)
	spin = spinner.New(charSet, spnTime, setOpts("Cloning repos", final))
	defer spin.Stop()
	spin.Start()

	for _, repo := range reposToClone {

		cloneURL := repo.CloneableURL(gls.Token)
		path := fmt.Sprintf("%s/%s", repo.Owner, repo.Name)

		go func(repoURL, dirPath string, w *sync.WaitGroup) {
			defer w.Done()
			err := CloneRepository(repoURL, dirPath)
			if err != nil {
				log.Println(err)
			}

		}(cloneURL, path, &wg)
	}
	wg.Wait()
}

// CloneRepository clones a repository
func CloneRepository(url, destDir string) (err error) {
	err = os.MkdirAll(destDir, 0755)
	if err != nil {
		return
	}
	_, err = git.PlainClone(destDir, false, &git.CloneOptions{URL: url})
	return
}

func setOpts(msg, final string) func(*spinner.Spinner) {
	return func(s *spinner.Spinner) {
		s.Suffix = fmt.Sprintf(" %s...", msg)
		s.FinalMSG = final
		s.Color("fgHiCyan")
		s.Writer = os.Stderr
		s.HideCursor = true
	}
}

func stats(owner string, repos []*repo) string {
	var owned []string
	var guest []string
	for _, repo := range repos {
		if owner == repo.Owner {
			owned = append(owned, repo.Name)
			continue
		}
		guest = append(guest, repo.Name)
	}

	return fmt.Sprintf(`
Secret Repos owned by %s: %d
Secret Repos owned by others: %d 
`, owner, len(owned), len(guest),
	)
}
