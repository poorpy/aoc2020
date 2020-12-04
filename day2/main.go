package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules, err := read("day2.input")
	if err != nil {
		panic(err)
	}

	count1, count2 := 0, 0
	for _, rule := range rules {
		if rule.isValid() {
			count1 += 1
		}

		if rule.otherValid() {
			count2 += 1
		}
	}
	fmt.Printf("Answer: %v Answer2: %v\n", count1, count2)
}

type Rule struct {
	lowerBound int
	upperBound int
	letter     rune
	passwd     string
}

func (r *Rule) isValid() bool {
	count := 0
	for _, l := range r.passwd {
		if l == r.letter {
			count += 1
		}
	}
	return count >= r.lowerBound && count <= r.upperBound
}

func (r *Rule) otherValid() bool {
	first := rune(r.passwd[r.lowerBound-1]) == r.letter
	second := rune(r.passwd[r.upperBound-1]) == r.letter
	return first != second
}

func NewRule(line string) (Rule, error) {
	split := strings.Split(line, " ")
	lowerUpper := strings.Split(split[0], "-")

	lower, err := strconv.Atoi(lowerUpper[0])
	if err != nil {
		return Rule{}, err
	}

	upper, err := strconv.Atoi(lowerUpper[1])
	if err != nil {
		return Rule{}, err
	}

	return Rule{
		lowerBound: lower,
		upperBound: upper,
		letter:     rune(split[1][0]), // we ignore : character at the end
		passwd:     split[2],
	}, nil

}

func read(filename string) ([]Rule, error) {
	var rules []Rule
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rule, err := NewRule(scanner.Text())
		if err != nil {
			return rules, err
		}

		rules = append(rules, rule)
	}

	return rules, nil
}
