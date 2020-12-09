package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	arr := read("day9.input")
	num, err := first(25, arr)
	if err != nil {
		fmt.Printf("Err: %v\n", err)
	}
	fmt.Printf("Answer: %v\n", num)
	ans, err := windowSearch(num, arr)
	if err != nil {
		fmt.Printf("Err: %v\n", err)
	}
	sort.Ints(ans)
	fmt.Printf("Answer: %v\n", ans[0]+ans[len(ans)-1])
}

func windowSearch(num int, arr []int) ([]int, error) {
	start, sum := 0, 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum == num {
			return arr[start : i+1], nil
		} else if sum > num {
			sum = 0
			start += 1
			i = start
		}
	}

	return []int{}, errors.New("No viable subarray found")
}

func first(pLen int, arr []int) (int, error) {
	tmp := make([]int, pLen)
	for i := pLen; i < len(arr); i++ {
		copy(tmp, arr[i-pLen:i])
		if !isSumOfPrevious(arr[i], tmp) {
			return arr[i], nil
		}
	}
	return 0, errors.New("No viable number found")
}

func isSumOfPrevious(num int, prev []int) bool {
	sort.Ints(prev)
	lower, upper := 0, len(prev)-1
	for lower < upper {
		test := prev[lower] + prev[upper]
		if test == num {
			return true
		} else if test < num {
			lower += 1
		} else if test > num {
			upper -= 1
		}
	}

	return false
}

func read(filename string) (ret []int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err == nil {
			ret = append(ret, n)
		}
	}

	return ret
}
