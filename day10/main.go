package main

import (
	"fmt"
	"strconv"
	"strings"
)

type machine struct {
	currentPosition int
	skipSize        int
	list            []int
	lengths         []int
}

func (m *machine) prefill() {
	m.skipSize = 0
	m.currentPosition = 0
	m.list = make([]int, 256)
	for idx := 0; idx < len(m.list); idx++ {
		m.list[idx] = idx
	}
}

func (m *machine) reverseList(length int) {
	for i := 0; i < length/2; i++ {
		frontIdx := (m.currentPosition + i) % len(m.list)
		backIdx := (m.currentPosition + length - 1 - i) % len(m.list)
		m.list[frontIdx], m.list[backIdx] = m.list[backIdx], m.list[frontIdx]
	}
}

func (m *machine) process() {
	for _, length := range m.lengths {
		m.reverseList(length)

		m.currentPosition += length + m.skipSize
		m.skipSize++
	}
}

func main() {
	lengths := "225,171,131,2,35,5,0,13,1,246,54,97,255,98,254,110"

	fmt.Println("Part 1:", part1(lengths))
	fmt.Println("Part 2:", part2(lengths))
}

func part2(lengths string) string {
	suffix := []int{17, 31, 73, 47, 23}
	splittedLengths := strings.Split(lengths, "")
	bytes := make([]int, len(splittedLengths))
	for i, l := range splittedLengths {
		bytes[i] = int([]byte(l)[0])
	}
	bytes = append(bytes, suffix...)

	m := machine{
		lengths: bytes,
	}
	m.prefill()

	for i := 0; i < 64; i++ {
		m.process()
	}

	return xor(m.list)
}

func part1(lengths string) string {
	splittedLengths := strings.Split(lengths, ",")
	bytes := make([]int, len(splittedLengths))
	for i, l := range splittedLengths {
		num, _ := strconv.Atoi(l)
		bytes[i] = num
	}

	m := machine{
		lengths: bytes,
	}
	m.prefill()

	m.process()

	return fmt.Sprintf("%v * %v = %v", m.list[0], m.list[1], m.list[0]*m.list[1])
}

func xor(s []int) string {
	var hexResult []string
	for i := 0; i < 256; i += 16 {
		xorResult := 0

		chunk := s[i : i+16]

		for _, num := range chunk {
			xorResult ^= num
		}
		if xorResult < 16 {
			hexResult = append(hexResult, fmt.Sprintf("0%x", xorResult))
		} else {
			hexResult = append(hexResult, fmt.Sprintf("%x", xorResult))
		}
	}

	return strings.Join(hexResult, "")
}
