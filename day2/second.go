package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("vim-go")

	file, err := os.Open("input")
	defer file.Close()

	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(file)

	var divisors []int

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		items := strings.SplitAfter(line, "\t")
		var numbers []int
		for _, item := range items {
			num, err := strconv.Atoi(strings.TrimSpace(item))
			if err != nil {
				fmt.Printf("Error converting %v to string", item)
				os.Exit(1)
			}
			numbers = append(numbers, num)
		}

	wholecalculation:
		for _, operandOne := range numbers {
			fmt.Println(operandOne)
			for _, operandTwo := range numbers {
				if operandOne != operandTwo && operandOne%operandTwo == 0 {
					divisors = append(divisors, operandOne/operandTwo)
					break wholecalculation
				}
			}
		}
	}

	var checksum int

	fmt.Println(divisors)
	for _, divs := range divisors {
		checksum += divs
	}

	fmt.Println("Checksum:", checksum)

}

func findMinimum(array []int) int {
	var number int = array[0]

	for i := 0; i < len(array); i++ {
		if array[i] < number {
			number = array[i]
		}
	}

	return number
}

func findMaximum(array []int) int {
	var number int = array[0]

	for i := 0; i < len(array); i++ {
		if number < array[i] {
			number = array[i]
		}
	}

	return number
}
