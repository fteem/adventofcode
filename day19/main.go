package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Coordinate struct {
	Y int
	X int
}

type Direction struct {
	Y int
	X int
}

func walk(in string) (string, int) {
	maze := strings.Split(strings.Trim(in, "\n"), "\n")

	pos := Coordinate{Y: 0, X: strings.IndexRune(maze[0], '|')}
	dir := Direction{Y: 1, X: 0}
	result := make([]byte, 0)
	count := 0
	for {
		var prev Coordinate
		prev, pos = pos, Coordinate{pos.Y + dir.Y, pos.X + dir.X}
		count++
		if pos.Y < 0 || pos.Y >= len(maze) || pos.X < 0 || pos.X >= len(maze[0]) {
			break
		}

		char := maze[pos.Y][pos.X]
		if char >= 'A' && char <= 'Z' {
			result = append(result, char)
		}

		if char == ' ' {
			break
		}

		if char == '+' {
			for _, nextDir := range []Direction{Direction{-1, 0}, Direction{0, -1}, Direction{0, 1}, Direction{1, 0}} {
				next := Coordinate{pos.Y + nextDir.Y, pos.X + nextDir.X}
				if next == prev {
					continue
				}

				if next.Y < 0 || next.Y >= len(maze) || next.X < 0 || next.X >= len(maze[0]) {
					continue
				}

				nextChar := maze[next.Y][next.X]
				if nextChar == ' ' {
					continue
				}
				dir = nextDir
			}
		}
	}
	return string(result), count
}

func part1(in string) string {
	letters, _ := walk(in)
	return letters
}

func part2(in string) int {
	_, count := walk(in)
	return count
}

func main() {
	checkForInputFile()
	b, err := ioutil.ReadFile(os.Args[1])
	checkError(err)

	str := string(b)

	fmt.Println("Part 1:", part1(str))
	fmt.Println("Part 2:", part2(str))
}

func checkForInputFile() {
	if len(os.Args) == 1 {
		checkError(errors.New("Missing input file as first argument. Exiting."))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
