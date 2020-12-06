package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	groups, err := read("day6.input")
	if err != nil {
		panic(err)
	}
	any, every := 0, 0
	for _, group := range groups {
		any += len(group.any())
		every += len(group.every())
	}
	fmt.Printf("Anybody: %v Everybody: %v\n", any, every)
}

type Group struct {
	Answers []string
}

func (g *Group) composeAnswers() string {
	return strings.Join(g.Answers, "")
}

func (g *Group) every() (ret string) {
	com := g.composeAnswers()
	stats := make(map[rune]int)
	for _, char := range com {
		stats[char] += 1
	}
	for key, count := range stats {
		if count == len(g.Answers) {
			ret += string(key)
		}
	}
	return
}

func (g *Group) any() (ret string) {
	com := g.composeAnswers()
	for _, c := range com {
		if !strings.ContainsRune(ret, c) {
			ret += string(c)
		}
	}

	return
}

func read(filename string) ([]Group, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var ret []Group
	strs := strings.Split(string(file), "\n\n")
	for _, str := range strs {
		str = strings.Trim(str, "\n")
		group := Group{
			Answers: strings.Split(str, "\n"),
		}
		ret = append(ret, group)
	}

	return ret, nil
}
