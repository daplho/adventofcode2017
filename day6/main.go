// http://adventofcode.com/2017/day/6
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var memoryBanks []int = []int{0, 2, 7, 0}
	var memoryBanks []int = []int{4, 10, 4, 1, 8, 4, 9, 14, 5, 1, 14, 15, 0, 15, 3, 5}

	i := 0
	states := make([][]int, 0)
	steps := 0

	for {
		maxIndex := findMaxIndex(memoryBanks)
		redistributed := redistribute(maxIndex, memoryBanks)
		steps = contains(redistributed, states)
		if steps > 0 {
			break
		}

		states = append(states, make([]int, len(redistributed)))
		copy(states[i], redistributed)
		i++
	}

	fmt.Printf("Total number of steps: '%d'\n", i+1)
	fmt.Printf("Steps since same state: '%d'\n", steps)
}

func contains(array []int, arrays [][]int) int {
	for i, currentArray := range arrays {
		if reflect.DeepEqual(array, currentArray) {
			return len(arrays) - i
		}
	}

	return 0
}

func findMaxIndex(memoryBanks []int) int {
	max := 0
	for i, memoryBank := range memoryBanks {
		if memoryBank > memoryBanks[max] {
			max = i
		}
	}

	return max
}

func redistribute(start int, memoryBanks []int) []int {
	valueToRedistribute := memoryBanks[start]
	memoryBanks[start] = 0
	j := (start + 1) % len(memoryBanks)
	for i := valueToRedistribute; i > 0; i-- {
		memoryBanks[j]++
		j = (j + 1) % len(memoryBanks)
	}

	return memoryBanks
}
