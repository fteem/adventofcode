package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Layer struct {
	dpt int // depth
	rng int // range
}

type Firewall struct {
	layers []Layer
	maxDpt int // max depth
}

func NewFirewall(data map[int]int) Firewall {
	maxDepth := findLastDepth(data)
	firewall := Firewall{
		maxDpt: maxDepth,
	}
	firewall.layers = make([]Layer, maxDepth+1)

	for depth, rng := range data {
		layer := Layer{
			dpt: depth,
			rng: rng,
		}
		firewall.layers[depth] = layer
	}

	return firewall
}

func main() {
	checkForInputFile()

	f, err := os.Open(os.Args[1])
	checkError(err)

	data := parseFile(f)
	firewall := NewFirewall(data)
	severity := part1(firewall)
	fmt.Println(severity)

	delays := part2(firewall)
	fmt.Println(delays)
}

func part2(firewall Firewall) int {
	var severities []int
	packet := 0
	delays := 0

	for {
		for {
			if packet > firewall.maxDpt {
				break
			}

			// Skip checks for layers without any range
			if firewall.layers[packet].rng == 0 {
				packet++
				continue
			} else {
				scannerPosition := packet % (2 * (firewall.layers[packet].rng - 1))
				// If the scanner is on position zero, it caught the packet!
				if scannerPosition == 0 {
					severities = append(severities, firewall.layers[packet].dpt*firewall.layers[packet].rng)

				}

				packet++
			}
			sum := 0
			for _, s := range severities {
				sum += s
			}
			if sum == 0 {
				break
			}
		}
		delays++
	}

	return delays

}

func part1(firewall Firewall) int {
	var severities []int
	packet := 0

	for {
		if packet > firewall.maxDpt {
			break
		}

		// Skip checks for layers without any range
		if firewall.layers[packet].rng == 0 {
			packet++
			continue
		} else {
			scannerPosition := packet % (2 * (firewall.layers[packet].rng - 1))
			// If the scanner is on position zero, it caught the packet!
			if scannerPosition == 0 {
				severities = append(severities, firewall.layers[packet].dpt*firewall.layers[packet].rng)

			}

			packet++
		}
	}

	sum := 0
	for _, s := range severities {
		sum += s
	}
	return sum
}

// Util functions

func parseFile(file *os.File) map[int]int {
	s := bufio.NewScanner(file)

	data := make(map[int]int)

	for s.Scan() {
		row := strings.Fields(s.Text())
		layer, _ := strconv.Atoi(strings.Trim(row[0], ":"))
		rng, _ := strconv.Atoi(strings.Trim(row[1], "\n"))

		data[layer] = rng
	}

	return data
}

func findLastDepth(d map[int]int) int {
	largest := 0

	for depth, _ := range d {
		if depth > largest {
			largest = depth
		}
	}

	return largest
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
