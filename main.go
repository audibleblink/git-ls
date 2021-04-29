package main

import (
	"fmt"
	"os"

	"github.com/audibleblink/git-ls/pkg/gitls"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	cli := gitls.NewClient(token)

	baseURL := os.Getenv("GITHUB_API_BASE_URL")
	if baseURL != "" {
		cli = gitls.NewEnterpriseClient(baseURL, token)
	}

	if len(os.Args) <= 1 {
		usage()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "repos":
		cli.Repos()
	case "collabs":
		cli.Collabs()
	case "gists":
		cli.Gists()
	case "user":
		cli.User()
	case "export":
		privateOnly := false
		cli.Plunder(privateOnly)
	case "plunder":
		privateOnly := true
		cli.Plunder(privateOnly)
	default:
		fmt.Println("Not Implemented")
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, `
Usage: git-ls <sub-cmd>

<user>		Inspect properties of the token owner
<gists>		See all gists, public and private, to which this token owner has access
<repos>		See all repos, public and private, to which this token owner has access
<collabs>	See other users' repos, public and private, to which this token owner has access
<export>	Clone all repos, public and private (check your HD size first!)
<plunder>	Clones all private repos the token can access with wreckless abandon`)
}
