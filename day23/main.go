package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	checkForInputFile()

	f, err := os.Open(os.Args[1])
	checkError(err)

	s := bufio.NewScanner(f)
	var commands []string
	var names []string
	var newValues []string

	for s.Scan() {
		row := strings.Fields(s.Text())
		command := strings.Trim(row[0], " ")
		name := strings.Trim(row[1], " ")
		value := strings.Trim(row[2], "\n")

		commands = append(commands, command)
		names = append(names, name)
		newValues = append(newValues, value)
	}

	fmt.Println("Part 1:", part1(commands, names, newValues))
	fmt.Println("Part 2:", part2())
}

func part2() int {
	b := 99*100 + 100000
	c := b + 17000
	d := 0
	f := 0
	h := 0

	for {
		f = 1
		d = 2
		for d != b {
			if b%d == 0 {
				f = 0
			}

			d++
		}

		if f == 0 {
			h++
		}
		if b == c {
			return h
		}
		b += 17
	}
}

func part1(commands []string, names []string, newValues []string) int {
	registries := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0}
	i := 0
	muls := 0
	for i < len(commands) {
		switch commands[i] {
		case "set":
			registries[names[i]] = retrieveFromRegistry(newValues[i], registries)
		case "sub":
			registries[names[i]] -= retrieveFromRegistry(newValues[i], registries)
		case "mul":
			registries[names[i]] *= retrieveFromRegistry(newValues[i], registries)
			muls++
		case "jnz":
			if val, exists := registries[names[i]]; exists {
				if val != 0 {
					converted, e := strconv.Atoi(newValues[i])
					checkError(e)

					i += converted
					continue
				}
			} else {
				v, _ := strconv.Atoi(names[i])
				if v != 0 {
					converted, e := strconv.Atoi(newValues[i])
					checkError(e)

					i += converted
					continue
				}

			}
		}
		i++
	}
	return muls
}

// Util funcs

func retrieveFromRegistry(newValue string, registries map[string]int) int {
	if val, exists := registries[newValue]; exists {
		return val
	} else {
		converted, _ := strconv.Atoi(newValue)
		return converted
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
func checkForInputFile() {
	if len(os.Args) == 1 {
		checkError(errors.New("Missing input file as first argument. Exiting."))
	}
}
