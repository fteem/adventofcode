package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Program struct {
	name     string
	parent   string
	weight   int
	children []string
}

func main() {
	data := readFile("input")
	fmt.Println(part1(data))
	fmt.Println(part2(data, part1(data)))
}

func part1(data map[string]Program) string {
	for k, v := range data {
		if v.parent == "" {
			return k
		}
	}
	return ""
}

func readFile(filename string) map[string]Program {
	data := make(map[string]Program, 0)

	file, _ := os.Open(filename)
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		row := strings.Fields(s.Text())

		var currentProgram Program
		for idx, value := range row {
			if idx == 0 {
				// Set name
				if v, exists := data[value]; exists {
					currentProgram = v
				} else {
					currentProgram.name = value
				}
			} else if idx == 1 {
				// Set weight
				currentProgram.weight, _ = strconv.Atoi(value[1 : len(value)-1])
			} else if idx > 2 {
				// Set children
				// Strip ,'s
				if value[len(value)-1] == ',' {
					value = value[:len(value)-1]
				}
				// Check if child is already parsed
				if v, exists := data[value]; exists {
					// Set Parent of Child
					v.parent = currentProgram.name
					data[value] = v
				} else {
					// Create Child with empty weight
					data[value] = Program{name: value, weight: 0, parent: currentProgram.name}
				}
				// Add to Children
				currentProgram.children = append(currentProgram.children, value)

			}
		}
		data[currentProgram.name] = currentProgram
	}
	return data
}

func sumWeight(data map[string]Program, root string) int {
	var total = 0
	for _, v := range data[root].children {
		total += sumWeight(data, v)
	}
	return total + data[root].weight
}

func isBalanced(data map[string]Program, root string) (string, int) {
	occurances := make(map[int]int)
	children := make(map[int]string)
	// Count occurances of totals
	for _, v := range data[root].children {
		occurances[sumWeight(data, v)]++
		children[sumWeight(data, v)] = v
	}
	// Get unbalanced name and the differnce
	var result, odd, normal = "", 0, 0
	for k, v := range occurances {
		if v == 1 {
			result = children[k]
			odd = k
		} else {
			normal = k
		}
	}
	return result, normal - odd
}

func part2(data map[string]Program, root string) int {
	if wrong, diff := isBalanced(data, root); wrong != "" {
		if part2(data, wrong) == 0 {
			return data[wrong].weight + diff
		}
		return part2(data, wrong)
	}
	return 0
}
