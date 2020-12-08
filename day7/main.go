package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules := read("day7.input")
	bagMap := make(map[uint32]Bag)
	for _, rule := range rules {
		addRuleToMap(rule, &bagMap)
	}

	allConnected := allConnected("shiny", "gold", &bagMap)
	goldContains := goldContains("shiny", "gold", &bagMap)
	fmt.Printf("Bag connections: %v gold bags contains: %v\n", len(allConnected), goldContains)

}

type Bag struct {
	Adjective   string
	Color       string
	Hash        uint32
	Connections []Connection
}

func (b Bag) String() string {
	return fmt.Sprintf("{a: %s c: %s h: %d c: %v}", b.Adjective, b.Color, b.Hash, b.Connections)
}

func (b *Bag) contains(hash uint32) bool {
	for _, conn := range b.Connections {
		if conn.Hash == hash {
			return true
		}
	}
	return false
}

type Connection struct {
	Weight uint32
	Hash   uint32
}

func (c Connection) String() string {
	return fmt.Sprintf("{w: %d h: %d}", c.Weight, c.Hash)
}

func hash(s string) uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		fmt.Printf("Panic hashing string: %s\n", s)
		panic(err)
	}
	return h.Sum32()
}

func allConnected(adjective, color string, bagMap *BagMap) []uint32 {
	target := hash(adjective + color)
	return bfs(target, bagMap)
}

func goldContains(adjective, color string, bagMap *BagMap) uint32 {
	root := hash(adjective + color)
	return dfs(root, bagMap)
}

type BagMap = map[uint32]Bag
type BagSet = map[uint32]bool

func bfs(target uint32, bagMap *BagMap) []uint32 {
	visited := make(BagSet)
	visited[target] = true
	queue, ret := []uint32{target}, []uint32{}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		for key, value := range *bagMap {
			if !visited[key] && value.contains(item) {
				visited[key] = true
				queue = append(queue, key)
				ret = append(ret, key)
			}
		}
	}

	return ret
}

func dfs(root uint32, bagMap *BagMap) uint32 {
	rBag := (*bagMap)[root]
	if len(rBag.Connections) == 0 {
		return 0 // we hit bottom
	}
	acc := uint32(0)
	for _, conn := range rBag.Connections {
		acc += conn.Weight + (conn.Weight * dfs(conn.Hash, bagMap))
	}
	return acc
}

func addRuleToMap(row string, bagMap *map[uint32]Bag) {
	arr := strings.Split(row, "contain")
	bagArgs := strings.Split(arr[0], " ")
	connArgs := strings.Split(strings.Trim(arr[1], " ."), ",")
	var connections []Connection
	if len(connArgs) != 1 || connArgs[0][:2] != "no" {
		for _, conn := range connArgs {
			parts := strings.Split(strings.Trim(conn, " "), " ")
			w, _ := strconv.Atoi(parts[0])
			weight := uint32(w)
			connections = append(connections, Connection{
				Weight: weight,
				Hash:   hash(parts[1] + parts[2]),
			})
		}
	}
	bag := Bag{
		Adjective:   bagArgs[0],
		Color:       bagArgs[1],
		Hash:        hash(bagArgs[0] + bagArgs[1]),
		Connections: connections,
	}
	(*bagMap)[bag.Hash] = bag
}

func read(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	var ret []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret
}
