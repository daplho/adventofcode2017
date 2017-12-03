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

	fmt.Printf("%v\n", matrix)
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
