package main

import "fmt"

func main() {
	input := 356
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func part1(input int) int {
	pointer := 0
	var buffer []int
	buffer = append(buffer, 0)
	for i := 1; i <= 2017; i++ {
		pointer = (pointer+input)%i + 1
		buffer = append(buffer, 0)
		copy(buffer[pointer+1:], buffer[pointer:])
		buffer[pointer] = i
	}
	return buffer[pointer+1]
}

func part2(input int) int {
	pointer := 0
	val := 0
	for i := 1; i <= 50000000; i++ {
		pointer = (pointer+input)%i + 1
		if pointer == 1 {
			val = i
		}
	}
	return val
}
