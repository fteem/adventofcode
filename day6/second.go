package main

import (
	"fmt"
	"reflect"
)

func main() {
	cycles := 0
	input := []int{10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6}
	final, producedConfigurations := first(input)

	for {
		var maxIdx = 0
		for i, bank := range final {
			if bank > final[maxIdx] {
				maxIdx = i
			}
		}
		maxValue := input[maxIdx]

		input[maxIdx] = 0
		for i := 1; i <= maxValue; i++ {
			input[(maxIdx+i)%len(input)]++
		}

		if reflect.DeepEqual(input, final) {
			fmt.Println(cycles + 1)
			return
		}

		cycles++
	}
}

func first(input []int) ([][]int, []int) {
	cycles := 0
	producedConfigurations := make([][]int, 0)

	for {
		// Find max value in iteration
		maxIdx := maxIndex(input)
		maxValue := input[maxIdx]

		// Clean the maximum value
		input[maxIdx] = 0

		// Increment all memory banks
		for i := 1; i <= maxValue; i++ {
			input[(maxIdx+i)%len(input)]++
		}

		// Check if the produced configs contain the latest generated one
		if contains(producedConfigurations, input) {
			cycles++
			break
		}

		// Copy the new confugration to the past ones
		inputCopy := make([]int, len(input))
		copy(inputCopy, input)
		producedConfigurations = append(producedConfigurations, inputCopy)

		// Increment counter
		cycles++
	}

	return input, producedConfigurations
}

func contains(configurations [][]int, candidate []int) bool {
	for _, configuration := range configurations {
		if reflect.DeepEqual(configuration, candidate) {
			return true
		}
	}

	return false
}

func maxIndex(s []int) int {
	idx := 0

	for i := 0; i < len(s); i++ {
		if s[idx] < s[i] {
			idx = i
		}
	}

	return idx
}
