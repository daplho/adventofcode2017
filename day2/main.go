// http://adventofcode.com/2017/day/2
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

	checksum, checksum2 := 0, 0
	for _, row := range matrix {
		checksum += maxMinDiff(row)
		checksum2 += divisionDiff(row)
	}

	fmt.Printf("Checksum: %d\n", checksum)
	fmt.Printf("Checksum2: %d\n", checksum2)
}

func maxMinDiff(row []int) int {
	if len(row) == 0 {
		return 0
	}
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

func divisionDiff(row []int) int {
	for i := 0; i < len(row); i++ {
		for j := i + 1; j < len(row); j++ {
			if row[i]%row[j] == 0 {
				return row[i] / row[j]
			} else if row[j]%row[i] == 0 {
				return row[j] / row[i]
			} else {
				continue
			}
		}
	}

	return 0
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
