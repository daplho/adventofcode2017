package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	matrix := readMatrix("./input.txt")

	checksum := 0
	for _, row := range matrix {
		checksum += maxMinDiff(row)
	}

	fmt.Printf("Checksum: %d\n", checksum)
}

func maxMinDiff(row []int) int {
	min, max := row[0], row[0]

	for _, num := range row {
		if num < min {
			min = num
		}

		if max < num {
			max = num
		}
	}

	return max - min
}

func readMatrix(filepath string) [][]int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matrix := make([][]int, 0)
	for scanner.Scan() {
		matrix = append(matrix, readRow(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return matrix
}

func readRow(line string) []int {
	numbers := make([]int, 0)
	stringNums := strings.Fields(line)
	for _, strNum := range stringNums {
		num, _ := strconv.Atoi(strNum)
		numbers = append(numbers, num)
	}

	return numbers
}
