package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type Node struct {
	name       int
	neighbours []*Node
}

func (n *Node) addEdge(neighbour *Node) {
	n.neighbours = append(n.neighbours, neighbour)
	neighbour.neighbours = append(neighbour.neighbours, n)
}

func (n *Node) visit(size int, visited map[int]struct{}) {
	_, alreadyVisited := visited[n.name]

	if alreadyVisited {
		return
	}

	visited[n.name] = struct{}{}
	size++
	for _, neighbour := range n.neighbours {
		neighbour.visit(size, visited)
	}
}

func (n *Node) graphSize() int {
	size := 0
	visited := make(map[int]struct{})
	n.visit(size, visited)

	return size
}

func newNode(name int) *Node {
	n := Node{
		name: name,
	}
	n.neighbours = make([]*Node, 0)
	return &n
}

func findOrCreateNode(nodes map[int]*Node, name int) *Node {
	n := nodes[name]
	if n == nil {
		n = newNode(name)
		nodes[name] = n
	}
	return n
}

func countGroups(nodes map[int]*Node) int {
	group := 0
	visited := make(map[int]struct{})
	for _, node := range nodes {
		_, alreadyVisited := visited[node.name]
		if !alreadyVisited {
			node.visit(visited)
			group += 1
		}
	}
	return group
}

func main() {
	if len(os.Args) == 1 {
		check(errors.New("Missing input file as argument."))
	}

	f, err := os.Open(os.Args[1])
	check(err)

	s := bufio.NewScanner(f)

	nodes := make(map[int]*Node)
	for s.Scan() {
		name, neighbours := parseLine(s)

		node := findOrCreateNode(nodes, name)
		for _, n := range neighbours {
			neighbour := findOrCreateNode(nodes, n)
			node.addEdge(neighbour)
		}
	}

	fmt.Println(nodes)
	fmt.Println(nodes[0].graphSize())
	fmt.Println(countGroups(nodes))
}

func parseLine(s *bufio.Scanner) (int, []int) {
	row := strings.Split(s.Text(), " <-> ")
	name, _ := strconv.Atoi(row[0])
	neighboursAsStrings := strings.Split(row[1], ",")
	neighbours := make([]int, len(neighboursAsStrings))

	for i, n := range neighboursAsStrings {
		converted, err := strconv.Atoi(strings.Trim(n, " "))
		check(err)

		neighbours[i] = converted
	}

	return name, neighbours
}
