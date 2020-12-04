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
	nums := read("day1.input")
	sort.Ints(nums)
	a1, b1, err := find(2020, nums)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Answer: %v\n", a1*b1)

	a2, b2, c2, err := find2(2020, nums)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Answer2: %v\n", a2*b2*c2)
}

func find(target int, nums []int) (int, int, error) {
	upper := len(nums) - 1
	lower := 0
	for lower < upper {
		test := nums[lower] + nums[upper]
		if test == target {
			return nums[lower], nums[upper], nil
		} else if test < target {
			lower += 1
		} else if test > target {
			upper -= 1
		}
	}
	return 0, 0, errors.New("target not found")
}

func find2(target int, nums []int) (int, int, int, error) {
	a, b, c, movedb := 0, 1, 2, false
	for c < len(nums) {
		test := nums[a] + nums[b] + nums[c]
		// fmt.Printf("test: %v a: %v b: %v c: %v")
		if test == target {
			return nums[a], nums[b], nums[c], nil
		}
		if test < target {
			c += 1
		}
		if test > target {
			if !movedb {
				b += 1
				c = b + 1
				movedb = true
			} else {
				a += 1
				b = a + 1
				c = b + 1
				movedb = false
			}
		}
	}

	return 0, 0, 0, errors.New("target not found")
}

func read(filename string) (nums []int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

	return
}
