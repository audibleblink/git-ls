package main

import (
	"fmt"
	"os"

	"github.com/audibleblink/git-ls/pkg/gitls"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	cli := gitls.NewClient(token)

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
