package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Waypoint struct {
	x int
	y int
}

func newWaypoint() *Waypoint {
	w := &Waypoint{
		x: 10,
		y: 1,
	}
	return w
}

func (w *Waypoint) stride(direction string, distance int) {
	if direction == "N" {
		w.y += distance
	} else if direction == "S" {
		w.y -= distance
	} else if direction == "E" {
		w.x += distance
	} else if direction == "W" {
		w.x -= distance
	}
}

func (w *Waypoint) turn(direction string, count int) {
	for i := 0; i < count; i++ {
		if direction == "R" {
			w.x, w.y = w.y, w.x*-1
		} else {
			w.x, w.y = w.y*-1, w.x
		}
	}
}

type Boat struct {
	x int
	y int
	w *Waypoint
}

func newBoat() Boat {
	boat := Boat{
		x: 0,
		y: 0,
		w: newWaypoint(),
	}
	return boat
}

func (b *Boat) move(command string) {
	if command[0] == 'R' || command[0] == 'L' {
		degree, _ := strconv.Atoi(command[1:])
		degree /= 90
		b.w.turn(command[:1], degree)
	} else if command[0] == 'F' {
		distance, _ := strconv.Atoi(command[1:])
		for i := 0; i < distance; i++ {
			b.x += b.w.x
			b.y += b.w.y
		}
	} else {
		distance, _ := strconv.Atoi(command[1:])
		b.w.stride(command[:1], distance)
	}
}

func read(filename string) (ret []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		ret = append(ret, scan.Text())
	}

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	commands, boat := read("day12.input"), newBoat()
	for _, command := range commands {
		boat.move(command)
	}

	fmt.Printf("x: %d y: %d sum: %d\n", boat.x, boat.y, abs(boat.x)+abs(boat.y))
}
