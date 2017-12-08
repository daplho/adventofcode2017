// http://adventofcode.com/2017/day/5
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	instructions := readFile("./input.txt")

	steps := findExit(instructions)
	fmt.Printf("Steps to exit: '%d'\n", steps)
}

func findExit(instructions []int) int {
	i, steps := 0, 0
	for {
		current := i

		if i < 0 || i >= len(instructions) {
			break
		}

		i = i + instructions[i]
		instructions[current]++
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
