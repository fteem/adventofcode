package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	operand  string
	operator boolOperator
	value    int
}

func (op Operation) evaluate(registry map[string]int) bool {
	return op.operator.evaluate(registry[op.operand], op.value)
}

type boolOperator struct {
	operator string
}

func (bo boolOperator) evaluate(value1 int, value2 int) bool {
	if bo.operator == ">" {
		return value1 > value2
	} else if bo.operator == "<" {
		return value1 < value2
	} else if bo.operator == ">=" {
		return value1 >= value2
	} else if bo.operator == "<=" {
		return value1 <= value2
	} else if bo.operator == "==" {
		return value1 == value2
	} else if bo.operator == "!=" {
		return value1 != value2
	}

	return false
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide input file as first argument. Exiting.")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	defer f.Close()

	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	// Registry
	registry := make(map[string]int)

	highestValueHeld := 0

	s := bufio.NewScanner(f)

	for s.Scan() {
		row := strings.Fields(s.Text())

		var currentRegistryName string
		var operation string
		var incrementValue int
		var op Operation

		for idx, item := range row {
			if idx == 0 {
				// Get register
				currentRegistryName = item
			} else if idx == 1 {
				// Get operation
				operation = item
			} else if idx == 2 {
				// Get value
				incrementValue, err = strconv.Atoi(item)
				if err != nil {
					fmt.Println("Conversion error:", err)
					os.Exit(1)
				}
			} else if idx == 3 {
				// Don't do anything with 'if'
			} else if idx >= 4 {
				// Condition
				value, err := strconv.Atoi(row[6])
				if err != nil {
					fmt.Println("Conversion error:", err)
					os.Exit(1)
				}

				op.operand = row[4]
				op.operator = boolOperator{row[5]}
				op.value = value
			}
		}

		if _, exists := registry[currentRegistryName]; !exists {
			registry[currentRegistryName] = 0
		}
		if op.evaluate(registry) {
			if operation == "inc" {
				registry[currentRegistryName] += incrementValue
			} else {
				registry[currentRegistryName] -= incrementValue
			}
			if registry[currentRegistryName] > highestValueHeld {
				highestValueHeld = registry[currentRegistryName]
			}
		}
	}

	fmt.Println("Largest registry:", findMax(registry))
	fmt.Println("Highest value held:", highestValueHeld)
}

func findMax(m map[string]int) int {
	largest := 0
	for _, v := range m {
		if largest < v {
			largest = v
		}
	}

	return largest
}
