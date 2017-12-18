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
		var value string
		if len(row) > 2 {
			value = strings.Trim(row[2], "\n")
		}
		commands = append(commands, command)
		names = append(names, name)
		newValues = append(newValues, value)
	}

	fmt.Println("Part 1:", part1(commands, names, newValues))
	fmt.Println("Part 2:", part2(commands, names, newValues))
}

func part2(commands []string, names []string, newValues []string) int {
	counter := 0
	queueA := make([]int, 0)
	queueB := make([]int, 0)
	runProgram(commands, names, newValues, queueA, queueB)
	runProgram(commands, names, newValues, queueB, queueA)

	return counter
}

func runProgram(commands []string, names []string, newValues []string, receiveQueue []int, sendQueue []int) {
	registries := make(map[string]int)
	i := 0
	for i < len(commands) {
		switch commands[i] {
		case "set":
			registries[names[i]] = retrieveFromRegistry(newValues[i], registries)
		case "add":
			registries[names[i]] += retrieveFromRegistry(newValues[i], registries)
		case "mul":
			registries[names[i]] *= retrieveFromRegistry(newValues[i], registries)
		case "mod":
			registries[names[i]] %= retrieveFromRegistry(newValues[i], registries)
		case "snd":
			v := retrieveFromRegistry(newValues[i], registries)
			sendQueue = append(sendQueue, v)
		case "jgz":
			if registries[names[i]] > 0 {
				converted, err := strconv.Atoi(newValues[i])
				checkError(err)

				i += converted
				continue
			}
		case "rcv":
			if len(receiveQueue) == 0 {
				continue
			}
			var val int
			val, receiveQueue = receiveQueue[0], receiveQueue[1:]
			fmt.Println(val)
		}
		i++
	}
}

func part1(commands []string, names []string, newValues []string) int {
	registries := make(map[string]int)
	i := 0
	played := 0
	for i < len(commands) {
		switch commands[i] {
		case "set":
			registries[names[i]] = retrieveFromRegistry(newValues[i], registries)
		case "add":
			registries[names[i]] += retrieveFromRegistry(newValues[i], registries)
		case "mul":
			registries[names[i]] *= retrieveFromRegistry(newValues[i], registries)
		case "mod":
			registries[names[i]] %= retrieveFromRegistry(newValues[i], registries)
		case "snd":
			played = registries[names[i]]
		case "jgz":
			if registries[names[i]] > 0 {
				converted, err := strconv.Atoi(newValues[i])
				checkError(err)

				i += converted
				continue
			}
		case "rcv":
			if registries[names[i]] > 0 {
				return played
			}
		}
		i++
	}
	return played
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
