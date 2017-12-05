// http://adventofcode.com/2017/day/3

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func main() {
	finalValue := 289326

	x, y := createSpiralPoint(float64(finalValue))
	fmt.Printf("--- Part One ---\n")
	fmt.Printf("%d %d\n", int(x), int(y))
	fmt.Printf("Manhattan Distance: %d\n", int(math.Abs(x)+math.Abs(y)))

	fmt.Printf("--- Part Two ---\n")
	partTwoAnswer := partTwo(finalValue)
	fmt.Printf("Part Two Answer: %d\n", partTwoAnswer)
}

func partTwo(finalValue int) int {
	grid := make(map[Point]int)

	for i := 1; ; i++ {
		x, y := createSpiralPoint(float64(i))
		value := adjacentSum(grid, int(x), int(y))
		if value > finalValue {
			return value
		}
		point := &Point{int(x), int(y)}
		grid[*point] = value
	}

	return -1
}

func adjacentSum(grid map[Point]int, x int, y int) int {
	if x == 0 && y == 0 {
		return 1
	}

	sum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			point := &Point{int(x) + i, int(y) + j}
			value, ok := grid[*point]
			if ok {
				sum += value
			}
		}
	}

	return sum
}

func createSpiralPoint(n float64) (float64, float64) {
	k := math.Ceil((math.Sqrt(n) - 1) / 2)
	t := 2*k + 1
	m := t * t
	t = t - 1

	if n >= m-t {
		return k - (m - n), -k
	} else {
		m = m - t
	}
	if n >= m-t {
		return -k, -k + (m - n)
	} else {
		m = m - t
	}
	if n >= m-t {
		return -k + (m - n), k
	} else {
		return k, k - (m - n - t)
	}
}
