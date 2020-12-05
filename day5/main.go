package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	seats, err := read("day5.input")
	if err != nil {
		panic(err)
	}

	ans, ids := 0, []int{}
	for _, seat := range seats {
		if id := seatId(seat); id > ans {
			ans = id
			ids = append(ids, id)
		} else {
			ids = append(ids, id)
		}
	}
	sort.Ints(ids)
	fmt.Printf("Answer: %v\n", ans)
	fmt.Printf("Missing: %v\n", findMissing(ids))
}

func findMissing(arr []int) []int {
	min := arr[0]
	missing := []int{}
	for i, val := 0, min; i < len(arr); i, val = i+1, val+1 {
		if arr[i] != val {
			missing = append(missing, val)
			val = arr[i]
		}
	}
	return missing
}

func seatId(data string) int {
	lower, upper := 0, 127
	left, right := 0, 7
	for _, c := range data {
		if c == 'F' {
			upper = (upper + lower) / 2
		}
		if c == 'B' {
			lower = (upper + lower + 1) / 2
		}
		if c == 'L' {
			right = (right + left) / 2
		}
		if c == 'R' {
			left = (right + left + 1) / 2
		}
	}
	return (lower * 8) + left
}

func read(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var ret []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret, nil
}
