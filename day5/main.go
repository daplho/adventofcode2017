// http://adventofcode.com/2017/day/5
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type AdjustOffset func(int) int

func main() {
	instructionsOne := readFile("./input.txt")
	instructionsTwo := make([]int, len(instructionsOne))
	copy(instructionsTwo, instructionsOne)

	stepsOne := findExit(instructionsOne, func(x int) int { return 1 })
	fmt.Printf("Steps to Exit One: '%d'\n", stepsOne)

	stepsTwo := findExit(instructionsTwo, func(x int) int {
		if x >= 3 {
			return -1
		}
		return 1
	})
	fmt.Printf("Steps to Exit Two: '%d'\n", stepsTwo)
}

func findExit(instructions []int, f AdjustOffset) int {
	i, steps := 0, 0
	for {
		current := i

		if i < 0 || i >= len(instructions) {
			break
		}

		i = i + instructions[i]
		instructions[current] += f(instructions[current])
		steps++
	}

	return steps
}

func readFile(filepath string) []int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]int, 0)
	for scanner.Scan() {
		lines = append(lines, readRow(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func readRow(line string) int {
	num, _ := strconv.Atoi(line)
	return num
}
