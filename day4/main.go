package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Missing filename as first argument.")
		os.Exit(1)
	}

	f := openFile(os.Args[1])
	defer f.Close()

	data := make([][]string, 0)
	s := bufio.NewScanner(f)

	for s.Scan() {
		row := strings.Fields(s.Text())

		data = append(data, row)
	}
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

func part2(data [][]string) int {
	numValid := 0

	for _, row := range data {
		valid := true
		used := make(map[string]bool)
		for _, word := range row {
			for k := range used {
				keySplit := strings.Split(k, "")
				sort.Strings(keySplit)
				wordSplit := strings.Split(word, "")
				sort.Strings(wordSplit)
				if reflect.DeepEqual(keySplit, wordSplit) {
					valid = false
					break
				}
			}

			if valid {
				used[word] = true
			} else {
				break
			}
		}
		if valid {
			numValid++
		}
	}

	return numValid
}

func part1(data [][]string) int {
	numValid := 0

	for _, row := range data {
		used := make(map[string]bool)
		valid := true
		for _, word := range row {
			if !used[word] {
				used[word] = true
			} else {
				valid = false
			}
		}

		if valid {
			numValid++
		}
	}

	return numValid
}

func openFile(filename string) *os.File {
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return f
}
