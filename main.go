package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Label struct {
	Name  string
	Color string
}

type Milestone struct {
	Name string
}

type Config struct {
	Repos      []string
	Labels     []Label
	Milestones []Milestone
}

func main() {
	var conf Config
	filecontents, err := ioutil.ReadFile("config.yml.example")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(filecontents, &conf)
	fmt.Printf("%#v\n", &conf)
}
