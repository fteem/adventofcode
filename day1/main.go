package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file := openInputFile()
	defer file.Close()

	var nextIndex int
	var matchingNumbers []int

	sum := 0

	input := parse(file)
	numbers := parseNumbers(input)

	for idx, num := range numbers {
		nextIndex = idx + (len(numbers) / 2)

		if nextIndex+1 >= len(numbers) {
			nextIndex = nextIndex - len(numbers) + 1
		}

		fmt.Println(nextIndex)
		fmt.Printf("Checking if %v is same as %v\n", num, numbers[nextIndex])
		if num == numbers[nextIndex] {
			fmt.Println("Ye, adding it do sum...")
			matchingNumbers = append(matchingNumbers, num)
			sum += num
		}
		fmt.Println("----------------------------")
	}
	fmt.Println(matchingNumbers)
	fmt.Println(sum)
}

func openInputFile() *os.File {
	file, err := os.Open("input")

	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}

	return file
}

func parse(f *os.File) string {
	reader := bufio.NewReader(f)
	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return line
}

func parseNumbers(s string) []int {
	var ints []int
	chars := []rune(s)
	for _, c := range chars {
		i, _ := strconv.Atoi(string(c))
		ints = append(ints, i)
	}
	return ints
}
