package main

import (
	"fmt"
	"github.com/google/go-github/github"
	oauth "golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
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

func getLabelsToAdd(client *github.Client, currentLabels, expectedLabels []github.Label) (labelsToAdd []github.Label) {
	labelsToAdd = make([]github.Label, 0) // have to make this 0 or else the first n elements are nil

	// turn  expectedLabels into a hash "set", for easy searching
	currentlabelset := make(map[string]bool)
	for _, l := range currentLabels {
		currentlabelset[l.String()] = false
	}

	for _, exlabel := range expectedLabels {
		if _, ok := currentlabelset[exlabel.String()]; !ok {
			labelsToAdd = append(labelsToAdd, exlabel)
		}
	}

	return
}

func getMilestonesToAdd(client *github.Client, currentMilestones, expectedMilestones []github.Milestone) (milestonesToAdd []github.Milestone) {
	milestonesToAdd = make([]github.Milestone, 0) // have to make this 0 or else the first n elements are nil

	// turn  expectedLabels into a hash "set", for easy searching
	currentmilestoneset := make(map[string]bool)
	for _, l := range currentMilestones {
		currentmilestoneset[l.String()] = false
	}

	for _, exlabel := range expectedMilestones {
		if _, ok := currentmilestoneset[exlabel.String()]; !ok {
			milestonesToAdd = append(milestonesToAdd, exlabel)
		}
	}

	return
}

func main() {
	var conf Config
	filecontents, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(filecontents, &conf)

	// retrieve information about the repos we're gonna manage
	client := setupClient()

	for _, repo := range conf.Repos {
		spl := strings.Split(repo, "/")
		owner := spl[0]
		reponame := spl[1]

		currentlabels, _, err := client.Issues.ListLabels(owner, reponame, nil)
		if err != nil {
			panic(err)
		}

		labelstoadd := getLabelsToAdd(client, currentlabels, conf.Labels)
		fmt.Printf("% v\n", labelstoadd)
	}
}
