package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	a *big.Int
	b *big.Int
}

func read(filename string) (int, []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timestamp, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	arr := strings.Split(scanner.Text(), ",")
	return timestamp, arr
}

func filterStops(names []string) (ret []int) {
	for _, ID := range names {
		if num, err := strconv.Atoi(ID); err == nil {
			ret = append(ret, num)
		}
	}
	return
}

func filterPairs(names []string) (ret []Pair) {
	for index, ID := range names {
		if num, err := strconv.ParseInt(ID, 10, 64); err == nil {
			ret = append(ret, Pair{a: big.NewInt(num - int64(index)), b: big.NewInt(num)})
		}
	}
	return
}

func crt(pairs []Pair) big.Int {
	M, total := big.NewInt(1), big.NewInt(0)
	for _, pair := range pairs {
		M.Mul(M, pair.b)
	}

	for _, pair := range pairs {
		b, tmp, exp := big.Int{}, big.Int{}, big.NewInt(-2)
		b.Div(M, pair.b)
		tmp.Mul(pair.a, &b)
		exp.Add(exp, pair.b).Exp(&b, exp, pair.b)
		tmp.Mul(&tmp, exp)
		total.Add(total, &tmp)
		total.Mod(total, M)
	}
	// My eyes are bleeding
	return *total
}

func timeToDeparture(timestamp, busID int) int {
	return busID - (timestamp % busID)
}

func main() {
	timestamp, buses := read("day13.input")
	busIDs := filterStops(buses)
	ID, delay := 0, 100
	for _, id := range busIDs {
		if diff := timeToDeparture(timestamp, id); diff < delay {
			ID, delay = id, diff
		}
	}
	fmt.Printf("Answer: %d\n", ID*delay)
	bigx := crt(filterPairs(buses))
	fmt.Printf("Answer2: %s\n", bigx.String())
}
