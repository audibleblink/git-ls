package gitls

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	// "github.com/google/go-github/v25/github"
)

type user struct {
	Bio                     string
	Company                 string
	Email                   string
	Location                string
	Login                   string
	Name                    string
	Type                    string
	Followers               int
	OwnedPrivateRepos       int
	PublicRepos             int
	PublicGists             int
	PrivateGists            int
	TotalPrivateRepos       int
	TwoFactorAuthentication bool
}

func (gls *ghClient) User() {
	usr, err := gls.user()
	if err != nil {
		log.Fatal(err)
	}
	out, err := json.MarshalIndent(usr, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(out))
}

func (gls *ghClient) user() (self *user, err error) {
	usr, _, err := gls.gh.Users.Get(context.Background(), "")
	if err != nil {
		return
	}
	self = &user{
		Bio:                     usr.GetBio(),
		Company:                 usr.GetCompany(),
		Email:                   usr.GetEmail(),
		Location:                usr.GetLocation(),
		Login:                   usr.GetLogin(),
		Name:                    usr.GetName(),
		Type:                    usr.GetType(),
		Followers:               usr.GetFollowers(),
		OwnedPrivateRepos:       usr.GetOwnedPrivateRepos(),
		PublicRepos:             usr.GetPublicRepos(),
		PublicGists:             usr.GetPublicGists(),
		PrivateGists:            usr.GetPrivateGists(),
		TotalPrivateRepos:       usr.GetTotalPrivateRepos(),
		TwoFactorAuthentication: usr.GetTwoFactorAuthentication(),
	}
	return
}
