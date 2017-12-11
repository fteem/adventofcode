package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(1)
	}
}

type hex struct {
	q float64
	r float64
	s float64
}

func (h *hex) move(step string) {
	switch step {
	case "n":
		h.q += 0
		h.r -= 1
	case "ne":
		h.q += 1
		h.r -= 1
	case "se":
		h.q += 1
		h.r += 0
	case "s":
		h.q += 0
		h.r += 1
	case "sw":
		h.q -= 1
		h.r += 1
	case "nw":
		h.q -= 1
		h.r += 0
	}
	h.s = -h.q - h.r
}

func NewHex() hex {
	return hex{
		q: 0.0,
		r: 0.0,
		s: 0.0,
	}
}

func distance(h1 hex, h2 hex) float64 {
	return math.Trunc((math.Abs(h1.q-h2.q) + math.Abs(h1.r-h2.r) + math.Abs(h1.s-h2.s)) / 2)
}

func parseFile(f *os.File) []string {
	r := bufio.NewReader(f)
	data, err := r.ReadString('\n')
	check(err)

	return strings.Split(strings.Trim(data, "\n"), ",")
}

func part1(steps []string) float64 {
	startingPoint := NewHex()
	h := NewHex()

	for _, step := range steps {
		h.move(step)
	}

	return distance(h, startingPoint)
}

func part2(steps []string) float64 {
	var maxDistance float64

	startingPoint := NewHex()
	h := NewHex()

	for _, step := range steps {
		h.move(step)
		currentDistance := distance(h, startingPoint)
		if maxDistance < currentDistance {
			maxDistance = currentDistance
		}
	}

	return maxDistance

}

func main() {
	if len(os.Args) == 1 {
		check(errors.New("Error: Missing input file as argument."))
	}

	f, err := os.Open(os.Args[1])
	check(err)
	defer f.Close()

	steps := parseFile(f)

	fmt.Println("Distance:", part1(steps))
	fmt.Println("Max distance:", part2(steps))
}
