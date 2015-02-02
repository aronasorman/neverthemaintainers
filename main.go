package main

import (
	"fmt"
	"github.com/google/go-github/github"
	oauth "golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Token string

func (t Token) Token() (*oauth.Token, error) {
	return &oauth.Token{AccessToken: string(t)}, nil
}

type Milestone struct {
	Name string
}

type Config struct {
	Repos      []string
	Labels     []github.Label
	Milestones []Milestone
}

func setupClient() *github.Client {
	token, _ := Token(os.Getenv("GITHUB_TOKEN")).Token()
	var conf oauth.Config
	client := conf.Client(nil, token)

	return github.NewClient(client)
}

func main() {
	var conf Config
	filecontents, err := ioutil.ReadFile("config.yml.example")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(filecontents, &conf)
	fmt.Printf("%#v\n", &conf)

	// retrieve information about the repos we're gonna manage
}
