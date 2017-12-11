package main

import "fmt"

func main() {
	input := 361527
	part1(input)
	part2(input)
}

func getNum(n, d int) int {
	return 4*n*n - (11-d)*n + (8 - d)
}

func part1(input int) {
	var n = 1
	for {
		var x, y int

		for d := 0; d < 7; d++ {
			if input >= getNum(n, d) && input < getNum(n, d+1) {
				switch d {
				case 0, 4:
					if input == getNum(n, d) {
						x, y = n-1, 0
					} else {
						x, y = n-1, input-getNum(n, d)
					}
				case 2, 6:
					if input == getNum(n, d) {
						x, y = 0, n-1
					} else {
						x, y = input-getNum(n, d), n-1
					}
				case 1, 3, 5:
					if input == getNum(n, d) {
						x, y = n-1, n-1
					} else {
						x, y = input-getNum(n, d), n-1
					}
				}
				fmt.Println("Part 1:", x+y)
				return
			}
		}
		if input >= getNum(n, 7) && input < getNum(n+1, 0) {
			if input == getNum(n, 7) {
				x = n
				y = n
			} else {
				if i := input - getNum(n, 7); i > 1 {
					x, y = n, i-1

				} else {
					x, y = n, n-1
				}
			}
			return
		}
		n++
	}
}

type cell struct {
	x int
	y int
}

func (c *cell) move(dir cell) cell {
	return cell{c.x + dir.x, c.y + dir.y}
}

func part2(input int) {
	store := map[cell]int{}
	value := 1
	actual := cell{0, 0}

	movements := []cell{
		cell{1, 0},
		cell{0, -1},
		cell{-1, 0},
		cell{0, 1},
	}

	neighbours := []cell{
		cell{1, 1},
		cell{-1, -1},
		cell{1, -1},
		cell{-1, 1},
	}
	neighbours = append(neighbours, movements...)

	store[actual] = value
	direction := 3

	for value <= input {
		newDirection := (direction + 1) % len(movements)
		candidate := actual.move(movements[newDirection])
		if _, ok := store[candidate]; ok {
			actual = actual.move(movements[direction])
		} else {
			actual = candidate
			direction = newDirection
		}
		value = 0
		for _, dir := range neighbours {
			neigh := actual.move(dir)
			val, ok := store[neigh]
			if ok {
				value = value + val
			}
		}
		store[actual] = value
	}
	fmt.Println("Result:", value)
}
