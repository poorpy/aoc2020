package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ret := read("day11.input")
	life := newLife(ret)
	for life.newStep() {
	}

	fmt.Printf("Answer: %v\n", life.a.totalOccupied())
}

func printBoard(board [][]rune) {
	for _, line := range board {
		for _, char := range line {
			fmt.Printf("%c ", char)
		}
		fmt.Println()
	}
}

type Field struct {
	b    [][]rune
	W, H int
}

func newField(board [][]rune) Field {
	return Field{
		b: board,
		W: len(board[0]),
		H: len(board),
	}
}

func (f *Field) totalOccupied() int {
	sum := 0
	for _, row := range f.b {
		for _, char := range row {
			if char == '#' {
				sum += 1
			}
		}
	}

	return sum
}

func (f *Field) nextState(x, y int) (rune, bool) {
	if f.b[x][y] == 'L' && f.areAdjacentEmpty(x, y) {
		return '#', true
	}
	if f.b[x][y] == '#' && f.countAdjacentTaken(x, y) >= 5 {
		return 'L', true
	}
	return f.b[x][y], false
}

func (f *Field) areAdjacentEmpty(x, y int) bool {
	areAdjacentEmpty := true
	for i := -1; i <= 1; i++ {
		if x+i <= -1 || x+i >= f.H {
			continue
		}
		for j := -1; j <= 1; j++ {
			if (y+j <= -1 || y+j >= f.W) || (i == 0 && j == 0) {
				continue
			}
			areAdjacentEmpty = areAdjacentEmpty && f.b[x+i][y+j] != '#'
		}
	}

	return areAdjacentEmpty
}

func (f *Field) countAdjacentTaken(x, y int) int {
	taken := 0
	for i := -1; i <= 1; i++ {
		if x+i <= -1 || x+i >= f.H {
			continue
		}
		for j := -1; j <= 1; j++ {
			if (y+j <= -1 || y+j >= f.W) || (j == 0 && i == 0) {
				continue
			}
			if f.b[x+i][y+j] == '#' {
				taken += 1
			}
		}
	}
	if x == 0 && y == 2 {
		fmt.Printf("taken: %d\n", taken)
	}

	return taken
}

type Life struct {
	a, b *Field
}

func newLife(board [][]rune) Life {
	field := newField(board)
	other := newField(make([][]rune, len(board)))
	for i := range other.b {
		other.b[i] = make([]rune, len(board[0]))
	}
	other.H, other.W = field.H, field.W
	return Life{
		a: &field,
		b: &other,
	}
}

func (l *Life) step() bool {
	stateChanged := false
	for x, row := range l.a.b {
		for y := range row {
			state, changed := l.a.nextState(x, y)
			l.b.b[x][y] = state
			stateChanged = stateChanged || changed
		}
	}

	l.a, l.b = l.b, l.a
	return stateChanged
}

func read(filename string) (board [][]rune) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner, board := bufio.NewScanner(file), [][]rune{}
	for scanner.Scan() {
		board = append(board, []rune(scanner.Text()))
	}

	return
}
