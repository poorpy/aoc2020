package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	slopeMap, err := read("day3.input")
	if err != nil {
		panic(err)
	}
	sum := count(slopeMap, 1, 1)
	sum *= count(slopeMap, 1, 3)
	sum *= count(slopeMap, 1, 5)
	sum *= count(slopeMap, 1, 7)
	sum *= count(slopeMap, 2, 1)
	fmt.Printf("Answer: %v\n", sum)
}

func count(slopeMap [][]rune, rowInc, columnInc int) int {
	rows := len(slopeMap)
	columns := len(slopeMap[0])
	count := 0
	for r, c := 0, 0; r < rows; r, c = r+rowInc, (c+columnInc)%columns {
		if slopeMap[r][c] == '#' {
			count += 1
		}
	}

	return count
}

func read(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var scanmap [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanmap = append(scanmap, []rune(scanner.Text()))
	}
	return scanmap, nil
}
