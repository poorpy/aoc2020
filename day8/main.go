package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instr struct {
	Opcode string
	Oper   int
}

func (i *Instr) swap() {
	if (*i).Opcode == "jmp" {
		(*i).Opcode = "nop"
	} else if (*i).Opcode == "nop" {
		(*i).Opcode = "jmp"
	}
}

func (i Instr) eval(index, acc *int) {
	if i.Opcode == "acc" {
		*acc += i.Oper
		*index += 1
	} else if i.Opcode == "jmp" {
		*index += i.Oper
	} else {
		*index += 1
	}
}

func main() {
	inst := read("day8.input")
	fmt.Printf("acc: %+v\n", evalOnce(inst))
	fmt.Printf("acc cor: %+v\n", evalCorrect(inst))
}

func evalOnce(boot []Instr) int {
	index, acc := 0, 0
	visited := make(map[int]bool)
	for {
		if !visited[index] {
			visited[index] = true
			boot[index].eval(&index, &acc)
		} else {
			break
		}
	}

	return acc
}

func testRun(boot []Instr) (int, bool) {
	index, acc, visited := 0, 0, make(map[int]bool)
	correct := true
	for index < len(boot) {
		if !visited[index] {
			visited[index] = true
			boot[index].eval(&index, &acc)
		} else {
			correct = false
			break
		}
	}

	return acc, correct
}

func evalCorrect(boot []Instr) int {
	toCheck := []int{}
	for i := range boot {
		if boot[i].Opcode == "jmp" || boot[i].Opcode == "nop" {
			toCheck = append(toCheck, i)
		}
	}

	for _, index := range toCheck {
		boot[index].swap()
		if acc, ok := testRun(boot); ok {
			return acc
		} else {
			boot[index].swap()
		}
	}

	panic("No instruction viable")
}

func read(filename string) []Instr {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var ret []Instr
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		i, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}
		ret = append(ret, Instr{
			Opcode: line[0],
			Oper:   i,
		})
	}

	return ret
}
