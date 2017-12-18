package main

import "fmt"

type Generator struct {
	factor int
	value  int
}

func (g Generator) newValue() int {
	return (g.value * g.factor) % 2147483647
}

func compare(valA int, valB int) bool {
	binA := fmt.Sprintf("%b", valA)
	if len(binA) < 16 {
		padding := 16 - len(binA)
		for i := 0; i < padding; i++ {
			binA = fmt.Sprintf("0%v", binA)
		}
	}
	binB := fmt.Sprintf("%b", valB)
	if len(binB) < 16 {
		padding := 16 - len(binB)
		for i := 0; i < padding; i++ {
			binB = fmt.Sprintf("0%v", binB)
		}
	}

	return binA[len(binA)-16:] == binB[len(binB)-16:]
}

func part1(gA Generator, gB Generator) int {
	count := 0
	for i := 1; i <= 40000000; i++ {
		if compare(gA.value, gB.value) {
			count++
		}
		gA.value = gA.newValue()
		gB.value = gB.newValue()
	}

	return count
}

func part2(gA Generator, gB Generator) int {
	count := 0
	valuesA := make([]int, 0)
	valuesB := make([]int, 0)

	for {
		if len(valuesA) == 5000000 {
			break
		}
		if gA.value%4 == 0 {
			valuesA = append(valuesA, gA.value)
		}
		gA.value = gA.newValue()
	}

	for {
		if len(valuesB) == 5000000 {
			break
		}
		if gB.value%8 == 0 {
			valuesB = append(valuesB, gB.value)
		}
		gB.value = gB.newValue()
	}

	for i := 0; i < 5000000; i++ {
		if compare(valuesA[i], valuesB[i]) {
			count++
		}
	}

	return count
}

func main() {
	generatorA := Generator{
		value:  289,
		factor: 16807,
	}
	generatorB := Generator{
		value:  629,
		factor: 48271,
	}

	fmt.Println("Part 1:", part1(generatorA, generatorB))
	fmt.Println("Part 2:", part2(generatorA, generatorB))
}
