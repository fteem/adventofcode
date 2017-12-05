package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := openInputFile()
	defer file.Close()

	reader := bufio.NewReader(file)

	var numbers []int

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		cleanLine := strings.TrimSpace(line)
		lineAsInt, _ := strconv.Atoi(cleanLine)
		numbers = append(numbers, lineAsInt)
	}

	fmt.Println(part1(numbers))
	fmt.Println(part2(numbers))
}

func part1(ints []int) int {
	steps := 0
	focusIndex := 0

	for {
		if len(ints) <= focusIndex {
			return steps
		}

		jump := ints[focusIndex]

		ints[focusIndex] += 1

		focusIndex += jump

		steps++
	}
}

func part2(ints []int) int {
	steps := 0
	focusIndex := 0

	for {
		if len(ints) <= focusIndex {
			return steps
		}

		jump := ints[focusIndex]

		if jump >= 3 {
			ints[focusIndex] -= 1
		} else {
			ints[focusIndex] += 1
		}

		focusIndex += jump

		steps++
	}

	return steps
}

func openInputFile() *os.File {
	file, err := os.Open("input")

	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}

	return file
}
