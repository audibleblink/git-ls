package main

import (
	"fmt"
	"os"

	"github.com/audibleblink/git-ls/pkg/gitls"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	cli := gitls.NewClient(token)

	if len(os.Args) <= 1 {
		usage()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "repos":
		cli.Repos()
	case "gists":
		cli.Gists()
	case "user":
		cli.User()
	default:
		fmt.Println("Not Implemented")
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, `
Usage: git-ls <repos | gists | user >

<user>
	Inspect properties of the token owner
<gists>
	See all gists, public and private, to which this token owner has access
<repos>
	See all repos, public and private, to which this token owner has access
	`)
}
