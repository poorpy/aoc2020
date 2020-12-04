package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var colorRegex *regexp.Regexp

func main() {
	arr, err := read("day4.input")
	if err != nil {
		panic(err)
	}
	colorRegex, _ = regexp.Compile(`^#([a-f]|\d){6}$`)
	if err != nil {
		panic(err)
	}

	var passports []Passport
	for _, pass := range arr {
		passports = append(passports, fromStrArr(pass))
	}
	sum := 0
	for _, pass := range passports {
		if pass.isValid2() {
			sum += 1
		}
	}
	fmt.Printf("Answer: %v\n", sum)
}

type Passport struct {
	Byr *string
	Iyr *string
	Eyr *string
	Hgt *string
	Hcl *string
	Ecl *string
	Pid *string
	Cid *string
}

func (p Passport) String() string {
	return fmt.Sprintf(
		"Byr: %s\nIyr: %s\nEyr: %s\nHgt: %s\nHcl: %s\nEcl: %s\nPid: %s\n",
		*p.Byr, *p.Iyr, *p.Eyr, *p.Hgt, *p.Hcl, *p.Ecl, *p.Pid,
	)
}

func (p *Passport) isValid() bool {
	return p.Byr != nil && p.Iyr != nil && p.Eyr != nil && p.Hgt != nil && p.Hcl != nil && p.Ecl != nil && p.Pid != nil
}

func (p *Passport) isValid2() bool {
	if p.isValid() {
		if i, err := strconv.Atoi(*p.Byr); len(*p.Byr) != 4 || err != nil || i < 1920 || i > 2002 {
			return false
		}

		if i, err := strconv.Atoi(*p.Iyr); len(*p.Iyr) != 4 || err != nil || i < 2010 || i > 2020 {
			return false
		}

		if i, err := strconv.Atoi(*p.Eyr); len(*p.Eyr) != 4 || err != nil || i < 2020 || i > 2030 {
			return false
		}

		hgt := *p.Hgt
		en := hgt[len(*p.Hgt)-2:]
		if en != "cm" && en != "in" {
			return false
		}

		if i, err := strconv.Atoi(hgt[:len(*p.Hgt)-2]); err != nil {
			return false
		} else if en == "cm" && (i < 150 || i > 193) {
			return false
		} else if en == "in" && (i < 59 || i > 76) {
			return false
		}

		hcl := []byte(*p.Hcl)
		if !colorRegex.Match(hcl) {
			return false
		}

		ecl, any := *p.Ecl, false
		for _, color := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
			if ecl == color {
				any = true
			}
		}
		if !any {
			return false
		}

		if _, err := strconv.Atoi(*p.Pid); err != nil || len(*p.Pid) != 9 {
			return false
		}

		return true

	}

	return false
}

func fromStrArr(arr []string) Passport {
	var pass Passport
	for _, str := range arr {
		str := strings.Split(str, ":")
		reflect.ValueOf(&pass).Elem().FieldByName(strings.Title(str[0])).Set(reflect.ValueOf(&str[1]))
	}
	return pass
}

func read(filename string) ([][]string, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var content [][]string
	strs := strings.Split(string(file), "\n\n")
	for _, str := range strs {
		str := strings.ReplaceAll(str, "\n", " ")
		str = strings.TrimRight(str, " ")
		arr := strings.Split(str, " ")
		content = append(content, arr)
	}

	return content, nil
}
