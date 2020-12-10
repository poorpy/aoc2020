package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	cache := make(map[int]int)
	arr := read("day10.input")
	ans := diff(arr)
	fmt.Printf("Answer: %v\n", ans)
	ans2 := combinations(0, &arr, &cache)
	fmt.Printf("Combinations: %v\n", ans2)
}

func combinations(index int, arr *[]int, cache *map[int]int) int {
	if index == len(*arr)-1 {
		return 1
	}
	if val, ok := (*cache)[index]; ok {
		return val
	}
	ans := 0
	for i := index + 1; i < len(*arr); i++ {
		if (*arr)[i]-(*arr)[index] <= 3 {
			ans += combinations(i, arr, cache)
		}
	}
	(*cache)[index] = ans
	return ans
}

func diff(arr []int) int {
	v1, v3 := 0, 0
	for i := 1; i < len(arr); i++ {
		test := arr[i] - arr[i-1]
		if test == 1 {
			v1 += 1
		} else if test == 3 {
			v3 += 1
		}
	}

	return v1 * v3
}

func read(filename string) (ret []int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		ret = append(ret, num)
	}
	ret = append([]int{0}, ret...)
	sort.Ints(ret)
	ret = append(ret, ret[len(ret)-1]+3)
	return ret
}
